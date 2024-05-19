@extends('admin.layout.welcome')
@section('title')
    Update Data Category
@endsection
@push('js')
    <script src="https://cdn.jsdelivr.net/npm/sweetalert2@11"></script>
    <script>
        function confirmUpdate(event) {
            event.preventDefault();
            Swal.fire({
                title: 'Are you sure?',
                text: 'You will update this data!',
                icon: 'warning',
                showCancelButton: true,
                confirmButtonColor: '#3085d6',
                cancelButtonColor: '#d33',
                confirmButtonText: 'Yes, update it!'
            }).then((result) => {
                if (result.isConfirmed) {
                    document.getElementById('update-form').submit();
                }
            });
        }
    </script>
@endpush
@section('content')
    <div class="container d-block">
        <form action="/admin/category/{{ $category['id'] }}" method="POST" id="update-form">
            @csrf
            @method('PUT')
            <div class="mb-3 pt-3">
                <label for="name" class="form-label">Category Name</label>
                <input type="text" class="form-control" name="name" value="{{ $category['name'] }}">
            </div>
            @error('name')
                <div class="alert alert-danger">{{ $message }}</div>
            @enderror
            <div class="mb-3">
                <label for="description" class="form-label">Category Description</label>
                <textarea name="description" id="" cols="30" rows="10" class="form-control">{{ $category['description'] }}</textarea>
            </div>
            @error('description')
                <div class="alert alert-danger">{{ $message }}</div>
            @enderror
            <button type="button" class="btn btn-primary" onclick="confirmUpdate(event)">Update</button>
        </form>
    </div>
@endsection
