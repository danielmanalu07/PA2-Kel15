<?php
namespace App\Http\Controllers\Admin;

use App\Http\Controllers\Controller;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Http;
use Illuminate\Support\Facades\Log;
use Illuminate\Support\Facades\Validator;

class ProductController extends Controller
{
    private $product = 'http://192.168.187.215:8080';
    private $category = 'http://192.168.187.215:8080';
    private $admin = 'http://192.168.187.215:8080/admin';

    public function index()
    {
        try {
            $token = session('jwt');

            $response = Http::withHeaders([
                'Cookie' => "jwt={$token}",
            ])->get("{$this->admin}/profile");

            $data = $response->json();

            $ProductResp = Http::get("{$this->product}/product");

            $products = $ProductResp->json();
            return view('admin.product.index', compact('data', 'products'));
        } catch (\Throwable $th) {
            throw $th;
        }
    }

    public function create()
    {
        try {
            $token = session('jwt');

            $response = Http::withHeaders([
                'Cookie' => "jwt={$token}",
            ])->get("{$this->admin}/profile");

            $CategoryResp = Http::get("{$this->category}/category");

            $category = $CategoryResp->json();

            $data = $response->json();

            return view('admin.product.create', compact('data', 'category'));
        } catch (\Throwable $th) {
            throw $th;
        }
    }

    public function store(Request $request)
    {
        $validator = Validator::make($request->all(), [
            'name' => 'required',
            'description' => 'required',
            'price' => 'required|numeric',
            'image' => 'required|image|mimes:jpeg,png,jpg|max:2048',
            'category_id' => 'required|integer',
        ]);

        if ($validator->fails()) {
            return redirect()->back()
                ->withErrors($validator)
                ->withInput();
        }

        try {
            $productResponse = Http::attach('image', $request->file('image')->get(), $request->file('image')->getClientOriginalName())
                ->post("{$this->product}/product/create", [
                    'name' => $request->name,
                    'description' => $request->description,
                    'price' => $request->price,
                    'category_id' => $request->category_id,
                ]);

            if ($productResponse->successful()) {
                return redirect('/admin/product')->with('success_message', 'Product created successfully!');
            } else {
                $errorMessage = $productResponse->json()['message'] ?? 'Failed to create product.';
                return back()->with('error_message', $errorMessage);
            }
        } catch (\Throwable $th) {
            throw $th;
        }
    }

    public function show($id)
    {
        try {
            $token = session('jwt');

            $response = Http::withHeaders([
                'Cookie' => "jwt={$token}",
            ])->get("{$this->admin}/profile");

            $data = $response->json();

            $productData = Http::get("{$this->product}/product/" . $id);

            if ($productData->successful()) {
                $responseArray = $productData->json();

                if (isset($responseArray['status']) && $responseArray['status'] === 'success' && isset($responseArray['message'])) {
                    $product = $responseArray['message'];

                    return view('admin.product.show', compact('product', 'data'));
                } else {
                    return redirect()->back()->with('error_message', 'Unexpected response structure. Please try again later.');
                }
            } else {
                return redirect()->back()->with('error_message', 'Failed to find product. Please try again later.');
            }
        } catch (\Throwable $th) {
            return redirect()->back()->with('error_message', 'Failed to find product. Please try again later.');
        }
    }

    public function edit($id)
    {
        try {
            $token = session('jwt');

            $response = Http::withHeaders([
                'Cookie' => "jwt={$token}",
            ])->get("{$this->admin}/profile");

            $data = $response->json();

            $productData = Http::get("{$this->product}/product/" . $id);
            // dd($productData->status(), $productData->body());

            if ($productData->successful()) {
                $responseArray = $productData->json();

                if (isset($responseArray['status']) && $responseArray['status'] === 'success' && isset($responseArray['message'])) {
                    $product = $responseArray['message'];

                    return view('admin.product.update', compact('product', 'data'));
                } else {
                    return redirect()->back()->with('error_message', 'Unexpected response structure. Please try again later.');
                }
            } else {
                return redirect()->back()->with('error_message', 'Failed to find product. Please try again later.');
            }
        } catch (\Throwable $th) {
            return redirect()->back()->with('error_message', 'Failed to find product. Please try again later.');
        }
    }

    public function update(Request $request, $id)
    {
        // Validate the incoming request
        $validator = Validator::make($request->all(), [
            'name' => 'required',
            'description' => 'required',
            'price' => 'required|numeric',
            'image' => 'nullable|image|mimes:jpeg,png,jpg|max:2048',
            'category_id' => 'required|integer',
        ]);

        // Handle validation failures
        if ($validator->fails()) {
            return redirect()->back()
                ->withErrors($validator)
                ->withInput();
        }

        try {
            // Prepare data for updating the product
            $updateData = [
                'name' => $request->name,
                'description' => $request->description,
                'price' => $request->price,
                'category_id' => $request->category_id,
            ];

            $token = session('jwt');

            // Check if an image file is included in the request
            if ($request->hasFile('image')) {
                // Attach the image file and send a POST request to update the product
                $productResponse = Http::withToken($token)
                    ->attach('image', $request->file('image')->get(), $request->file('image')->getClientOriginalName())
                    ->post("{$this->product}/product/edit/{$id}", $updateData);
            } else {
                // Send a PUT request without an image to update the product
                $productResponse = Http::withToken($token)
                    ->put("{$this->product}/product/edit/{$id}", $updateData);
            }

            // Check if the response indicates a successful update
            if ($productResponse->successful()) {
                return redirect('/admin/product')->with('success_message', 'Product updated successfully!');
            } else {
                // Extract and display error message from the response body
                $errorMessage = $productResponse->json()['message'] ?? 'Failed to update product.';
                $responseBody = $productResponse->body();
                return back()->with('error_message', $errorMessage . ' Response: ' . $responseBody);
            }
        } catch (\Throwable $th) {
            // Log the error for debugging purposes
            Log::error('Error updating product: ' . $th->getMessage());
            return back()->with('error_message', 'An unexpected error occurred. Please try again later.');
        }
    }

    public function destroy(string $id)
    {
        try {
            $token = session('jwt');

            $response = Http::delete("{$this->product}/product/delete/" . $id);

            $data = $response->json();

            if ($response->successful()) {
                return redirect()->back()->with('success_message', 'Deleted  Successfully!');
            } else {
                return redirect()->back()->with('error_message', 'Failed to delete product. Please try again later.');
            }
        } catch (\Throwable $th) {
            return redirect()->back()->with('error_message', 'Failed to delete product. Please try again later.');
        }
    }
}
