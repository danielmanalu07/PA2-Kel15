@extends('admin.layout.welcome')

@section('title')
    Edit Product
@endsection

@push('js')
    <script src="https://cdn.jsdelivr.net/npm/sweetalert2@11"></script>
    <script>
        function confirmUpdate(event) {
            event.preventDefault();
            Swal.fire({
                title: 'Are you sure?',
                text: 'You will update this product!',
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
    <div class="container mt-5">
        <div class="card shadow-lg">
            <div class="card-header bg-primary text-white">
                <h1 class="display-4" style="font-family: Lucida">{{ $product['name'] }}</h1>
            </div>
            <div class="card-body">
                <form action="/admin/product/{{ $product['id'] }}" method="POST" id="update-form" enctype="multipart/form-data">
                    @csrf
                    @method('PUT')
                    <div class="mb-3">
                        <label for="name" class="form-label">Product Name</label>
                        <input type="text" class="form-control" id="name" name="name" value="{{ $product['name'] }}">
                    </div>
                    @error('name')
                        <div class="alert alert-danger">{{ $message }}</div>
                    @enderror
                    <div class="mb-3">
                        <label for="description" class="form-label">Product Description</label>
                        <textarea class="form-control" id="description" name="description" rows="3">{{ $product['description'] }}</textarea>
                    </div>
                    @error('description')
                        <div class="alert alert-danger">{{ $message }}</div>
                    @enderror
                    <div class="mb-3">
                        <label for="price" class="form-label">Product Price</label>
                        <input type="number" class="form-control" id="price" name="price" value="{{ $product['price'] }}">
                    </div>
                    @error('price')
                        <div class="alert alert-danger">{{ $message }}</div>
                    @enderror
                    <div class="mb-3">
                        <label for="image" class="form-label">Product Image</label>
                        <input type="file" class="form-control" id="image" name="image">
                        <img src="http://192.168.100.24:8080/product/image/{{ $product['image'] }}" alt="Product Image" class="img-fluid mt-3" style="width: 30%; height: auto;">
                    </div>
                    @error('image')
                        <div class="alert alert-danger">{{ $message }}</div>
                    @enderror
                    <button type="button" class="btn btn-primary" onclick="confirmUpdate(event)">Update</button>
                    <a href="/admin/product" class="btn btn-secondary">Back to List</a>
                </form>
            </div>
        </div>
    </div>
@endsection
