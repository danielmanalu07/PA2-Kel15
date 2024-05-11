<?php

namespace App\Http\Controllers\Admin;

use App\Http\Controllers\Controller;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Http;
use Illuminate\Support\Facades\Validator;

class ProductController extends Controller
{
    private $product = 'http://127.0.0.1:8003';
    private $category = 'http://127.0.0.1:8002';
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
                ->post("http://127.0.0.1:8003/product/create", [
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
    public function show(string $id)
    {
        //
    }

    /**
     * Show the form for editing the specified resource.
     */
    public function edit(string $id)
    {
        //
    }

    /**
     * Update the specified resource in storage.
     */
    public function update(Request $request, string $id)
    {
        //
    }

    /**
     * Remove the specified resource from storage.
     */
    public function destroy(string $id)
    {
        //
    }
}
