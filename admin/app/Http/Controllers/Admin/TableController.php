<?php

namespace App\Http\Controllers\Admin;

use App\Http\Controllers\Controller;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Http;
use Illuminate\Support\Facades\Log;
use Illuminate\Support\Facades\Validator;

class TableController extends Controller
{
    private $tableService = 'http://172.26.43.150:8080';
    private $admin = 'http://172.26.43.150:8080/admin';

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
            // dd($tables); // Debug the table data
    
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
            // 'status' => 'required|string', // Remove this line
        ]);

        if ($validator->fails()) {
            return redirect()->back()
                ->withErrors($validator)
                ->withInput();
        }

        try {
            $status = $request->input('status', 'kosong'); // Set default status to "pending"

            $tableResp = Http::post("{$this->tableService}/table/create", [
                'number' => intval($request->input('number')),
                'capacity' => intval($request->input('capacity')),
                'status' => $status,
            ]);

            if ($tableResp->successful()) {
                return redirect('/admin/table')->with('success_message', 'Table created successfully!');
            } else {
                $errorMessage = $tableResp->json()['message'] ?? 'Failed to create table.';
                return back()->with('error_message', $errorMessage);
            }
        } catch (\Throwable $th) {
            Log::error('Table creation failed: ' . $th->getMessage());
            return back()->with('error_message', 'Failed to create table. Please try again later.');
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

            $tableData = Http::get("{$this->tableService}/table/" . $id);

            if ($tableData->successful()) {
                $responseArray = $tableData->json();

                if (isset($responseArray['status']) && $responseArray['status'] === 'success' && isset($responseArray['message'])) {
                    $table = $responseArray['message'];

                    return view('admin.table.show', compact('table', 'data'));
                } else {
                    return redirect()->back()->with('error_message', 'Unexpected response structure. Please try again later.');
                }
            } else {
                return redirect()->back()->with('error_message', 'Failed to find table. Please try again later.');
            }
        } catch (\Throwable $th) {
            return redirect()->back()->with('error_message', 'Failed to find table. Please try again later.');
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

            $tableData = Http::get("{$this->tableService}/table/" . $id);

            if ($tableData->successful()) {
                $responseArray = $tableData->json();

                if (isset($responseArray['status']) && $responseArray['status'] === 'success' && isset($responseArray['message'])) {
                    $table = $responseArray['message'];

                    return view('admin.table.update', compact('table', 'data'));
                } else {
                    return redirect()->back()->with('error_message', 'Unexpected response structure. Please try again later.');
                }
            } else {
                return redirect()->back()->with('error_message', 'Failed to find table. Please try again later.');
            }
        } catch (\Throwable $th) {
            return redirect()->back()->with('error_message', 'Failed to find table. Please try again later.');
        }
    }

    /**
     * Update the specified resource in storage.
     */
    public function update(Request $request, string $id)
    {
        $validator = Validator::make($request->all(), [
            'number' => 'required|integer',
            'capacity' => 'required|integer',
            'status' => 'required|string',
        ]);

        if ($validator->fails()) {
            return redirect()->back()->withErrors($validator)->withInput();
        }

        try {
            $token = session('jwt');

            $response = Http::put("{$this->tableService}/table/edit/" . $id, [
                'number' => intval($request->input('number')),
                'capacity' => intval($request->input('capacity')),
                'status' => $request->input('status'),
            ]);

            if ($response->successful()) {
                return redirect('/admin/table')->with('success_message', 'Table updated successfully!');
            } else {
                return redirect()->back()->with('error_message', 'Failed to update table. Please try again later.');
            }
        } catch (\Throwable $th) {
            return redirect()->back()->with('error_message', 'Failed to update table. Please try again later.');
        }
    }

    /**
     * Remove the specified resource from storage.
     */
    public function destroy(string $id)
    {
        try {
            $response = Http::delete("{$this->tableService}/table/delete/" . $id);

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
