@extends('admin.layout.welcome')
@section('title')
Show Table
@endsection

@section('content')
<div class="container">
    <div class="row mb-3">
        <div class="col-md-12">
            <h2>Table Details</h2>
        </div>
    </div>
    @if (Session::has('error_message'))
    <div class="alert alert-danger alert-dismissible fade show" role="alert">
        <strong>Error: </strong> {{ Session::get('error_message') }}
        <button type="button" class="close" data-dismiss="alert" aria-label="Close">
            <span aria-hidden="true">&times;</span>
        </button>
    </div>
    @endif
    <div class="card">
        <div class="card-header">
            Table ID: {{ $table['id'] }}
        </div>
        <div class="card-body">
            <h5 class="card-title">Table Number: {{ $table['number'] }}</h5>
            <p class="card-text">Capacity: {{ $table['capacity'] }}</p>
            {{-- <p class="card-text">Status: {{ $table['status'] }}</p> --}}
            <a href="/admin/table/{{ $table['id'] }}/edit" class="btn btn-warning"><i class="fas fa-edit"></i> Edit</a>
            <a href="/admin/table" class="btn btn-secondary"><i class="fas fa-arrow-left"></i> Back to List</a>
        </div>
    </div>
</div>
@endsection
