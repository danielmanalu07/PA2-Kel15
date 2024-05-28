<?php

namespace App\Http\Controllers\Admin;

use App\Http\Controllers\Controller;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Http;
use Illuminate\Support\Facades\Log;

class AdminController extends Controller
{
    private $apiUrl = 'http://172.26.43.150:8080/admin';

    public function adminLogin(Request $request)
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
                    // Log the response to get more details
                    Log::error('Login API failed', ['response' => $response->body()]);
                    return redirect()->back()->with('error', 'Invalid credentials or server error.');
                }

                $data = $response->json();
                session(['jwt' => $data['token']]);
                return redirect('/admin/dashboard')->with('success', 'Login successfully.');

            } catch (\Throwable $th) {
                // Log the exception to get more details
                Log::error('Login failed', ['exception' => $th]);
                return redirect()->back()->with('error', 'Internal Server Error.');
            }
        }

        return view('admin.login');
    }

    public function getProfile(Request $request)
    {
        try {
            $token = session('jwt');
            if (!$token) {
                return redirect('/admin/login')->with('error', 'Unauthenticated');
            }

            $response = Http::withHeaders([
                'Cookie' => "jwt={$token}",
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

    public function logoutAdmin(Request $request)
    {
        try {
            $token = session('jwt');
            if (!$token) {
                return redirect('/admin/login')->with('error', 'Unauthenticated');
            }
    
            $response = Http::withHeaders([
                'Cookie' => "jwt={$token}",
            ])->post("{$this->apiUrl}/logout");
    
            session()->forget('jwt');
            return redirect('/admin/login')->with('success', 'Logout successful');
    
        } catch (\Throwable $th) {
            return redirect('/admin/login')->with('error', 'You must be logged in.');
        }
    }
    

    public function profile(Request $request)
    {
        try {
            $token = session('jwt');
            if (!$token) {
                return redirect('/admin/login')->with('error', 'Unauthenticated');
            }

            $response = Http::withHeaders([
                'Authorization' => "Bearer {$token}",
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

    // Add the dashboard method here

    public function dashboard(Request $request)
        {
            try {
                $token = session('jwt');
                if (!$token) {
                    return redirect('/admin/login')->with('error', 'Unauthenticated');
                }
                $response = Http::withHeaders([
                    'Cookie' => "jwt={$token}",
                ])->get("{$this->apiUrl}/profile");
    
                if ($response->failed()) {
                    throw new \Exception("Failed to fetch profile.");
                }
    
                $data = $response->json();
    
                return view('admin.dashboard', compact('data')); // Ensure you have a 'dashboard.blade.php' file in the 'admin' view directory
            } catch (\Throwable $th) {
                return redirect('/admin/login')->with('error', 'You must be logged in.');
            }
        }

    
    public function approve(string $id)
    {
        try {
            $token = session('jwt');
    
            $response = Http::put("http://172.26.43.150:8080/order/status/" . $id, [
                'status' => 1,
            ]);
    
            if ($response->successful()) {
                return redirect('/admin/order')->with('success_message', 'Order approved successfully!');
            } else {
                return redirect()->back()->with('error_message', 'Failed to update order. Please try again later.');
            }
        } catch (\Throwable $th) {
            return redirect()->back()->with('error_message', 'Failed to update order. Please try again later.');
        }
    }
    
    public function reject(string $id)
    {
        try {
            $token = session('jwt');
    
            $response = Http::put("http://172.26.43.150:8080/order/status/" . $id, [
                'status' => 2,
            ]);
    
            if ($response->successful()) {
                return redirect('/admin/order')->with('success_message', 'Order rejected successfully!');
            } else {
                return redirect()->back()->with('error_message', 'Failed to update order. Please try again later.');
            }
        } catch (\Throwable $th) {
            return redirect()->back()->with('error_message', 'Failed to update order. Please try again later.');
        }
    }
        
}
