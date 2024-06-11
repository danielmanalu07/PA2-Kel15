<?php

namespace App\Http\Controllers\Admin;

use App\Http\Controllers\Controller;
use Illuminate\Support\Facades\Http;
use Illuminate\Support\Facades\Log;

class OrderController extends Controller
{
    private $orderService = 'http://192.168.187.215:8080';
    private $admin = 'http://192.168.187.215:8080/admin';

    /**
     * Display a listing of the resource.
     */
    public function index()
    {
        try {
            $token = session('jwt');

            // Fetch the admin profile
            $response = Http::withHeaders([
                'Cookie' => "jwt={$token}",
            ])->get("{$this->admin}/profile");

            $data = $response->json();

            // Fetch the list of orders
            $orderResp = Http::get("{$this->orderService}/order");

            if ($orderResp->successful()) {
                $responseData = $orderResp->json();

                // Check if the 'message' key exists in the response data
                if (isset($responseData['message'])) {
                    $orders = $responseData['message']; // Access the 'message' key
                    return view('admin.order.index', compact('data', 'orders'));
                } else {
                    return back()->with('error_message', 'No orders found.');
                }
            } else {
                $errorMessage = $orderResp->json()['message'] ?? 'Failed to fetch orders.';
                return back()->with('error_message', $errorMessage);
            }
        } catch (\Throwable $th) {
            Log::error('Order fetching failed: ' . $th->getMessage());
            return back()->with('error_message', 'Failed to fetch orders. Please try again later.');
        }
    }

    /**
     * Display the specified resource.
     */
    public function show($id)
    {
        try {
            $token = session('jwt');

            $response = Http::withHeaders([
                'Cookie' => "jwt={$token}",
            ])->get("{$this->admin}/profile");

            $data = $response->json();

            // Convert $id to string to avoid type issues
            $orderData = Http::get("{$this->orderService}/order/" . strval($id));

            if ($orderData->successful()) {
                $order = $orderData->json();
                return view('admin.order.show', compact('order', 'data'));
            } else {
                return redirect()->back()->with('error_message', 'Failed to find order. Please try again later.');
            }
        } catch (\Throwable $th) {
            return redirect()->back()->with('error_message', 'Failed to find order. Please try again later.');
        }
    }
}
