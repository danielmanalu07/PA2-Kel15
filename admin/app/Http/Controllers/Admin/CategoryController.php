<?php

namespace App\Http\Controllers\Admin;

use App\Http\Controllers\Controller;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Http;
use Illuminate\Support\Facades\Validator;

class CategoryController extends Controller
{
    private $apiUrl = 'http://127.0.0.1:8002';
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

            $CategoryResp = Http::get("{$this->apiUrl}/category");

            $category = $CategoryResp->json();

            return view('admin.category.index', compact('data', 'category'));

            // return response()->json($category);

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

            $data = $response->json();

            return view('admin.category.create', compact('data'));
        } catch (\Throwable $th) {
            throw $th;
        }
    }

    /**
     * Store a newly created resource in storage.
     */
    public function store(Request $request)
    {
        $validator = Validator::make($request->all(), [
            'name' => 'required',
            'description' => 'required',
        ]);
        if ($validator->fails()) {
            return redirect()->back()->withErrors($validator)->withInput();
        }
        try {

            $response = Http::post("{$this->apiUrl}/category/create", [
                'name' => $request->name,
                'description' => $request->description,
            ]);

            $data = $response->json();

            if ($response->successful()) {
                return redirect('/admin/category')->with('success_message', 'Created  Successfully!');
            } else {
                return redirect()->back()->with('error_message', 'Failed to create category. Please try again later.');
            }

        } catch (\Throwable $th) {
            return redirect()->back()->with('error_message', 'Failed to create category. Please try again later.');
        }
    }

    /**
     * Display the specified resource.
     */
    public function show(string $id)
    {
        try {
            $token = session('jwt');

            $response = Http::withHeaders([
                'Cookie' => "jwt={$token}",
            ])->get("{$this->admin}/profile");

            $data = $response->json();

            $categoryData = Http::get("{$this->apiUrl}/category/" . $id);

            $category = $categoryData->json();

            if ($categoryData->successful()) {
                return view('admin.category.show', compact('category', 'data'));
            } else {
                return redirect()->back()->with('error_message', 'Failed to find category. Please try again later.');
            }

        } catch (\Throwable $th) {
            return redirect()->back()->with('error_message', 'Failed to find category. Please try again later.');
        }
    }

    /**
     * Show the form for editing the specified resource.
     */
    public function edit(string $id)
    {
        try {
            $token = session('jwt');

            $response = Http::withHeaders([
                'Cookie' => "jwt={$token}",
            ])->get("{$this->admin}/profile");

            $data = $response->json();

            $categoryData = Http::get("{$this->apiUrl}/category/" . $id);

            $category = $categoryData->json();

            if ($categoryData->successful()) {
                return view('admin.category.update', compact('category', 'data'));
            }
        } catch (\Throwable $th) {
            //throw $th;
            return redirect()->back()->with('error_message', 'Failed to find category. Please try again later.');
        }
    }

    /**
     * Update the specified resource in storage.
     */
    public function update(Request $request, string $id)
    {
        $validator = Validator::make($request->all(), [
            'name' => 'required',
            'description' => 'required',
        ]);
        if ($validator->fails()) {
            return redirect()->back()->withErrors($validator)->withInput();
        }

        try {
            $token = session('jwt');

            $response = Http::put("{$this->apiUrl}/category/" . $id . "/edit", [
                'name' => $request->input('name'),
                'description' => $request->input('description'),
            ]);

            $data = $response->json();

            $CategoryResp = Http::get("{$this->apiUrl}/category/");

            $category = $CategoryResp->json();

            if ($CategoryResp->successful()) {
                return redirect('/admin/category')->with('success_message', 'Updated  Successfully!');
            } else {
                return redirect()->back()->with('error_message', 'Failed to update category. Please try again later.');
            }

            // return dd($data);

        } catch (\Throwable $th) {
            //throw $th;
            return redirect()->back()->with('error_message', 'Failed to update category. Please try again later.');

        }
    }

    /**
     * Remove the specified resource from storage.
     */
    public function destroy(string $id)
    {
        try {
            $token = session('jwt');

            $response = Http::delete("{$this->apiUrl}/category/" . $id . "/delete");

            $data = $response->json();

            // return dd($data);
            if ($response->successful()) {
                return redirect()->back()->with('success_message', 'Deleted  Successfully!');
            } else {
                return redirect()->back()->with('error_message', 'Failed to delete category. Please try again later.');
            }
        } catch (\Throwable $th) {
            return redirect()->back()->with('error_message', 'Failed to delete category. Please try again later.');
        }
    }
}
