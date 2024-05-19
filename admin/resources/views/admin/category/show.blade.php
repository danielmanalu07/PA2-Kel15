@extends('admin.layout.welcome')

@section('title')
    Detail Category Data
@endsection

@section('content')
    <div class="container mt-5">
        <div class="card shadow-lg">
            <div class="card-header bg-primary text-white">
                <h1 class="display-4" style="font-family: Lucida">{{ $category['name'] }}</h1>
            </div>
            <div class="card-body">
                <p class="lead" style="font-family: 'Comic Sans MS', cursive, sans-serif">{{ $category['description'] }}</p>
                <a href="/admin/category" class="btn btn-info btn-sm mt-3">Back to Category List</a>
            </div>
        </div>
    </div>
@endsection
