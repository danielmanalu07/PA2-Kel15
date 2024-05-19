@extends('admin.layout.welcome')

@section('title')
    Create Data Product
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
                    // Submit the form if confirmed
                    document.getElementById('create-form').submit();
                }
            });
        }

        function validateForm(event) {
            // Optionally, you can add more validation logic here
            confirmCreate(event);
        }
    </script>
@endpush

@section('content')
    <div class="container">
        @if (Session::has('error_message'))
            <div class="alert alert-danger alert-dismissible fade show" role="alert">
                <strong>Message: </strong> {{ Session::get('error_message') }}
                <button type="button" class="close" data-dismiss="alert" aria-label="Close">
                    <span aria-hidden="true">&times;</span>
                </button>
            </div>
        @endif
        <form action="{{ url('/admin/product') }}" method="POST" enctype="multipart/form-data" id="create-form">
            @csrf
            <div class="mb-3">
                <label for="name" class="form-label">Product Name <span style="color: red">*</span></label>
                <input type="text" class="form-control" id="name" name="name" value="{{ old('name') }}">
                @error('name')
                    <div class="alert alert-danger mt-1">{{ $message }}</div>
                @enderror
            </div>
            <div class="mb-3">
                <label for="description" class="form-label">Product Description <span style="color: red">*</span></label>
                <textarea class="form-control" id="description" name="description">{{ old('description') }}</textarea>
                @error('description')
                    <div class="alert alert-danger mt-1">{{ $message }}</div>
                @enderror
            </div>
            <div class="mb-3">
                <label for="image" class="form-label">Product Image <span style="color: red">*</span></label>
                <input type="file" class="form-control" id="image" name="image">
                @error('image')
                    <div class="alert alert-danger mt-1">{{ $message }}</div>
                @enderror
            </div>
            <div class="mb-3">
                <label for="price" class="form-label">Product Price <span style="color: red">*</span></label>
                <input type="number" class="form-control" id="price" name="price" value="{{ old('price') }}">
                @error('price')
                    <div class="alert alert-danger mt-1">{{ $message }}</div>
                @enderror
            </div>
            <div class="mb-3">
                <label for="category_id" class="form-label">Category <span style="color: red">*</span></label>
                <select class="form-control" id="category_id" name="category_id">
                    <option value="">--- Select Category ---</option>
                    @foreach ($category['message'] as $cat)
                        <option value="{{ $cat['id'] }}" {{ old('category_id') == $cat['id'] ? 'selected' : '' }}>
                            {{ $cat['name'] }}
                        </option>
                    @endforeach
                </select>
                @error('category_id')
                    <div class="alert alert-danger mt-1">{{ $message }}</div>
                @enderror
            </div>
            <button type="button" class="btn btn-primary" onclick="validateForm(event)">Create</button>
        </form>
    </div>
@endsection
