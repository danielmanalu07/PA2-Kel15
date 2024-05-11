<?php

namespace App\Http\Controllers\Admin;

use App\Http\Controllers\Controller;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Http;

class AdminController extends Controller
{
    private $apiUrl = 'http://127.0.0.1:8080/admin';

    public function login(Request $request)
    {
        if ($request->isMethod('post')) {
            $username = $request->input('username');
            $password = $request->input('password');
            if (empty($username) || empty($password)) {
                return redirect()->back()->with('error', 'Username and password are required.');
            }

            try {
                $response = Http::post("{$this->apiUrl}/login", [
                    'username' => $username,
                    'password' => $password,
                ]);

                if ($response->failed()) {
                    return redirect()->back()->with('error', 'Invalid credentials or server error.');
                }

                $data = $response->json();
                session(['jwt' => $data['token']]);
                return redirect('/admin/dashboard')->with('success', 'Login successfully.');

            } catch (\Throwable $th) {
                return redirect()->back()->with('error', 'Internal Server Error.');
            }
        }

        return view('admin.login');
    }

    public function dashboard(Request $request)
    {
        try {
            $token = session('jwt');
            $response = Http::withHeaders([
                'Cookie' => "jwt={$token}",
            ])->get("{$this->apiUrl}/profile");

            if ($response->failed()) {
                throw new \Exception("Failed to fetch profile.");
            }

            $data = $response->json();
            return view('admin.dashboard', compact('data'));
            // return response()->json($data);

        } catch (\Throwable $th) {
            return redirect('/admin/login')->with('error', 'You must be logged in.');
        }
    }

    public function logout()
    {
        try {
            $token = session('admin');
            $response = Http::withHeaders([
                'Cookie' => "admin={$token}",
            ])->post("{$this->apiUrl}/logout");

            session()->forget('admin');
            return redirect('/admin/login')->with('success', 'Logout successful');

        } catch (\Throwable $th) {
            return redirect('/admin/login')->with('error', 'You must be logged in.');
        }
    }

    public function profile()
    {
        try {
            $token = session('admin');
            $response = Http::withHeaders([
                'Cookie' => "admin={$token}",
            ])->get("{$this->apiUrl}/profile");

            if ($response->failed()) {
                throw new \Exception("Failed to fetch profile.");
            }

            $data = $response->json();
            return view('admin.account.profile', compact('data'));

        } catch (\Throwable $th) {
            return redirect('/admin/login')->with('error', 'You must be logged in.');
        }
    }
}
