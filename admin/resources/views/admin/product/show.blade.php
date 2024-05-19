@extends('admin.layout.welcome')

@section('title')
    Product Details
@endsection

@section('content')
    <div class="container mt-5">
        <div class="card shadow-lg">
            <div class="card-header bg-primary text-white">
                <h1 class="display-4" style="font-family: Lucida">{{ $product['name'] }}</h1>
            </div>
            <div class="card-body">
                <div class="mb-3">
                    <img src="http://127.0.0.1:8080/product/image/{{ $product['image'] }}" alt="Product Image"
                        class="img-fluid rounded mb-3" style="max-width: 100%; height: auto;">
                </div>
                <div class="mb-3">
                    <h5 class="card-title">Description</h5>
                    <p class="card-text">{{ $product['description'] }}</p>
                </div>
                <div class="mb-3">
                    <h5 class="card-title">Price</h5>
                    <p class="card-text">Rp. {{ $product['price'] }}</p>
                </div>
                <a href="/admin/product" class="btn btn-info btn-sm">Back to List</a>
            </div>
        </div>
    </div>
@endsection
