<?php

namespace App\Http\Controllers\Admin;

use App\Http\Controllers\Controller;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Http;
use Illuminate\Support\Facades\Validator;

class ProductController extends Controller
{
    private $product = 'http://127.0.0.1:8080';
    private $category = 'http://127.0.0.1:8080';
    private $admin = 'http://127.0.0.1:8080/admin';
    /**
     * Display a listing of the resource.
     */
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

    /**
     * Show the form for creating a new resource.
     */
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


    /**
     * Show the form for editing the specified resource.
     */
    public function edit($id)
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

    public function update(Request $request, string $id)
{
    $validator = Validator::make($request->all(), [
        'name' => 'required|string|max:255',
        'description' => 'required|string',
        'price' => 'required|numeric',
        'image' => 'nullable|image|mimes:jpeg,png,jpg,gif|max:2048',
    ]);

    if ($validator->fails()) {
        return redirect()->back()->withErrors($validator)->withInput();
    }

    try {
        $token = session('jwt');

        $response = Http::withHeaders([
            'Cookie' => "jwt={$token}",
        ])->get("{$this->admin}/profile");

        $data = $response->json();

        // Fetch the existing product to get the current image
        $productResponse = Http::get("{$this->product}/product/" . $id);
        if (!$productResponse->successful()) {
            return redirect()->back()->with('error_message', 'Failed to retrieve product data. Please try again later.');
        }
        $product = $productResponse->json()['message'];

        // Prepare form data
        $formData = [
            'name' => $request->input('name'),
            'description' => $request->input('description'),
            'price' => $request->input('price'),
            'image' => $product['image'], // Use existing image as default
        ];

        if ($request->hasFile('image')) {
            $imagePath = $request->file('image')->store('product_images', 'public');
            $formData['image'] = $imagePath;
        }

        // Update the product
        $updateResponse = Http::put("{$this->product}/product/" . $id . "/edit", $formData);

        if ($updateResponse->successful()) {
            return redirect('/admin/product')->with('success_message', 'Product updated successfully.');
        } else {
            return redirect()->back()->with('error_message', 'Failed to update product. Please try again later.');
        }
    } catch (\Throwable $th) {
        return redirect()->back()->with('error_message', 'Failed to update product. Please try again later.');
    }
}


    /**
     * Remove the specified resource from storage.
     */
    public function destroy(string $id)
    {
        //
    }
}
