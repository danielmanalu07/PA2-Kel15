<?php

namespace App\Http\Controllers\Admin;

use App\Http\Controllers\Controller;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Http;
use Illuminate\Support\Facades\Log;

class AdminController extends Controller
{
    private $apiUrl = 'http://172.27.1.162:8080/admin';
    private $requestTable = 'http://172.27.1.162:8080/requestTable';

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

    public function approveReqTable(Request $request, $id)
    {
        try {
            $token = session('jwt');
            if (!$token) {
                return redirect('/admin/login')->with('error', 'Unauthenticated');
            }

            // Convert id to integer
            $id = intval($id);
            if ($id <= 0) {
                return redirect()->back()->with('error', 'Invalid table ID.');
            }

            // Parse input from request
            $input = $request->all();

            // Call the API to approve the request table
            $response = Http::withHeaders([
                'Cookie' => "jwt {$token}",
            ])->put("{$this->apiUrl}/table/{$id}", $input);

            if ($response->failed()) {
                return redirect()->back()->with('error', 'Failed to approve request table. ' . $response->body());
            }

            return redirect()->back()->with('success', 'Request table approved successfully.');
        } catch (\Throwable $th) {
            Log::error('Failed to approve request table', ['exception' => $th]);
            return redirect()->back()->with('error', 'Failed to approve request table.');
        }
    }

    public function getAllRequestTables(Request $request)
    {
        try {
            $token = session('jwt');
            if (!$token) {
                return redirect('/admin/login')->with('error', 'Unauthenticated');
            }

            // Request for admin data
            $adminResponse = Http::withHeaders([
                'Cookie' => "jwt={$token}",
            ])->get("{$this->apiUrl}/profile");

            if ($adminResponse->failed()) {
                dd('Failed to fetch admin data', $adminResponse->body()); // Debugging line
                Log::error('Failed to fetch admin data', ['response' => $adminResponse->body()]);
                return redirect()->back()->with('error', 'Failed to fetch admin data. Please check the logs for more details.');
            }

            // Request for request tables data
            $requestTableResponse = Http::withHeaders([
                'Cookie' => "jwt={$token}",
            ])->get("{$this->requestTable}");

            if ($requestTableResponse->failed()) {
                dd('Failed to fetch request tables', $requestTableResponse->body()); // Debugging line
                Log::error('Failed to fetch request tables', ['response' => $requestTableResponse->body()]);
                return redirect()->back()->with('error', 'Failed to fetch request tables. Please check the logs for more details.');
            }

            $data = $adminResponse->json();
            $requestTables = $requestTableResponse->json();

            // Debugging line to check the response data
            // dd('Admin Data:', $adminData, 'Request Tables:', $requestTables);

            return view('admin.requestTable.index', compact('data', 'requestTables'));
        } catch (\Throwable $th) {
            Log::error('Failed to fetch data', ['exception' => $th]);
            return redirect()->back()->with('error', 'Failed to fetch data.');
        }
    }


    public function approveRequestTable(Request $request, $id)
    {
        return $this->updateRequestTableStatus($request, $id, 1);
    }

    public function rejectRequestTable(Request $request, $id)
    {
        return $this->updateRequestTableStatus($request, $id, 2);
    }

    public function pendingRequestTable(Request $request, $id)
    {
        return $this->updateRequestTableStatus($request, $id, 0);
    }

    private function updateRequestTableStatus(Request $request, $id, $status)
{
    try {
        $token = session('jwt');
        if (!$token) {
            return redirect('/admin/login')->with('error', 'Unauthenticated');
        }

        $id = intval($id);
        if ($id <= 0) {
            return redirect()->back()->with('error', 'Invalid request table ID.');
        }

        Log::info('Updating request table status', ['id' => $id, 'status' => $status]);

        $response = Http::withHeaders([
            'Cookie' => "jwt={$token}",
        ])->put("{$this->apiUrl}/table/{$id}", ['status' => $status]);

        if ($response->failed()) {
            Log::error('Failed to update request table status', ['response' => $response->body()]);
            return redirect()->back()->with('error', 'Failed to update request table status. ' . $response->body());
        }

        return redirect()->back()->with('success', 'Request table status updated successfully.');
    } catch (\Throwable $th) {
        Log::error('Failed to update request table status', ['exception' => $th]);
        return redirect()->back()->with('error', 'Failed to update request table status.');
    }
}

}
