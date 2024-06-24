@extends('admin.layout.welcome')

@section('title', 'Update Data Tabel')

@section('content')
<div class="container">
    <div class="row mb-3">
        <div class="col-md-6">
            <h2>Edit Table</h2>
        </div>
    </div>

    @if (Session::has('success_message'))
    <div class="alert alert-success alert-dismissible fade show" role="alert">
        <strong>Success:</strong> {{ Session::get('success_message') }}
        <button type="button" class="close" data-dismiss="alert" aria-label="Close">
            <span aria-hidden="true">&times;</span>
        </button>
    </div>
    @endif

    @if (Session::has('error_message'))
    <div class="alert alert-danger alert-dismissible fade show" role="alert">
        <strong>Error:</strong> {{ Session::get('error_message') }}
        <button type="button" class="close" data-dismiss="alert" aria-label="Close">
            <span aria-hidden="true">&times;</span>
        </button>
    </div>
    @endif

    <form action="{{ route('admin.table.update', $table['id']) }}" method="POST">
        @csrf
        @method('PUT')
        <div class="form-group">
            <label for="number">Number</label>
            <input type="text" name="number" class="form-control" id="number" value="{{ old('number', $table['number']) }}" required>
        </div>
        <div class="form-group">
            <label for="capacity">Capacity</label>
            <input type="number" name="capacity" class="form-control" id="capacity" value="{{ old('capacity', $table['capacity']) }}" required>
        </div>
        <div class="form-group">
            <label for="status">Status</label>
            <select name="status" class="form-control" id="status" required>
                <option value="kosong" {{ old('status', $table['status']) == 'kosong' ? 'selected' : '' }}>Tersedia</option>
                <option value="penuh" {{ old('status', $table['status']) == 'penuh' ? 'selected' : '' }}>Sedang digunakan</option>
            </select>
        </div>
        <button type="submit" class="btn btn-primary">Update</button>
    </form>
</div>
@endsection
