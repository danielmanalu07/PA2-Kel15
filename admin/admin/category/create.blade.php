@extends('admin.layout.welcome')
@section('title')
    Create Data Category
@endsection
@push('js')
    <script src="https://cdn.jsdelivr.net/npm/sweetalert2@11"></script>
    <script>
        function confirmCreate(event) {
            event.preventDefault();
            Swal.fire({
                title: 'Are you sure?',
                text: 'You will create this data!',
                icon: 'warning',
                showCancelButton: true,
                confirmButtonColor: '#d33',
                cancelButtonColor: '#3085d6',
                confirmButtonText: 'Yes, create it!'
            }).then((result) => {
                if (result.isConfirmed) {
                    document.getElementById('create-form').submit();
                }
            });
        }

        function validateForm(event) {
            confirmCreate(event);
        }
    </script>
@endpush
@section('content')
    <div class="container d-block">
        <form action="/admin/category" method="POST" id="create-form">
            @csrf
            <div class="mb-3 pt-3">
                <label for="name" class="form-label">Category Name</label>
                <input type="text" class="form-control" id="name" name="name">
            </div>
            @error('name')
                <div class="alert alert-danger">{{ $message }}</div>
            @enderror
            <div class="mb-3">
                <label for="description" class="form-label">Category Description</label>
                <textarea name="description" id="description" cols="30" rows="10" class="form-control"></textarea>
            </div>
            @error('description')
                <div class="alert alert-danger">{{ $message }}</div>
            @enderror
            <button type="button" class="btn btn-primary" onclick="validateForm(event)">Create</button>
        </form>
    </div>
@endsection
