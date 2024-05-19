<?php

namespace App\Http\Controllers\Admin;

use App\Http\Controllers\Controller;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Http;
use Illuminate\Support\Facades\Validator;

class TableController extends Controller
{
    private $tableService = 'http://127.0.0.1:8080';
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
            $tableResp = Http::get("{$this->tableService}/table");

            $tables = $tableResp->json();
            return view('admin.table.index', compact('data', 'tables'));
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

            $tableResp = Http::get("{$this->tableService}/table");

            $category = $tableResp->json();

            $data = $response->json();

            return view('admin.table.create', compact('data'));
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
            'number' => 'required|integer',
            'capacity' => 'required|integer',
        ]);

        if ($validator->fails()) {
            return redirect()->back()
                ->withErrors($validator)
                ->withInput();
        }

        try {
            $TableResp = Http::post("{$this->tableService}/table/create", [
                'number' => $request->number,
                'capacity' => $request->capacity,
            ]);

            if ($TableResp->successful()) {
                return redirect('/admin/table')->with('success_message', 'Table created successfully!');
            } else {
                $errorMessage = $TableResp->json()['message'] ?? 'Failed to create table.';
                return back()->with('error_message', $errorMessage);
            }
        } catch (\Throwable $th) {
            return back()->with('error_message', 'Failed to create table. Please try again later.');
        }
    }


    /**
     * Display the specified resource.
     */
    public function show(string $id)
    {
        try {
            $response = Http::get("{$this->tableService}/{$id}");

            if ($response->successful()) {
                $table = $response->json();
                return view('admin.table.show', compact('table'));
            } else {
                return redirect('/admin/table')->with('error_message', 'Table not found.');
            }
        } catch (\Throwable $th) {
            throw $th;
        }
    }

    /**
     * Show the form for editing the specified resource.
     */
    // public function edit(string $id)
    // {
    //     try {
    //         $response = Http::get("{$this->tableService}/{$id}");

    //         if ($response->successful()) {
    //             $table = $response->json();
    //             return view('admin.table.edit', compact('table'));
    //         } else {
    //             return redirect('/admin/table')->with('error_message', 'Table not found.');
    //         }
    //     } catch (\Throwable $th) {
    //         throw $th;
    //     }
    // }

    /**
     * Update the specified resource in storage.
     */
    public function update(Request $request, string $id)
    {
        $validator = Validator::make($request->all(), [
            'number' => 'required|integer',
            'capacity' => 'required|integer',
        ]);

        if ($validator->fails()) {
            return redirect()->back()->withErrors($validator)->withInput();
        }

        try {
            $response = Http::put("{$this->tableService}/{$id}/edit", [
                'number' => $request->number,
                'capacity' => $request->capacity,
            ]);

            if ($response->successful()) {
                return redirect('/admin/table')->with('success_message', 'Table updated successfully!');
            } else {
                return redirect()->back()->with('error_message', 'Failed to update table.');
            }
        } catch (\Throwable $th) {
            return redirect()->back()->with('error_message', 'Failed to update table.');
        }
    }

    /**
     * Remove the specified resource from storage.
     */
    public function destroy(string $id)
    {
        try {
            $response = Http::delete("{$this->tableService}/{$id}/delete");

            if ($response->successful()) {
                return redirect()->back()->with('success_message', 'Table deleted successfully!');
            } else {
                return redirect()->back()->with('error_message', 'Failed to delete table.');
            }
        } catch (\Throwable $th) {
            return redirect()->back()->with('error_message', 'Failed to delete table.');
        }
    }
}
