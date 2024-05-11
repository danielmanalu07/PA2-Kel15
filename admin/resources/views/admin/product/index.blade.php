@extends('admin.layout.welcome')
@section('title')
    List Product
@endsection
@section('content')
    <div class="container">
        <div class="row mb-3">
            <div class="col-md-3">
                <input type="text" id="search" class="form-control" placeholder="Search...">
            </div>
        </div>
        @if (Session::has('success_message'))
            <div class="alert alert-success alert-dismissible fade show" role="alert">
                <strong>Success: </strong> {{ Session::get('success_message') }}
                <button type="button" class="close" data-dismiss="alert" aria-label="Close">
                    <span aria-hidden="true">&times;</span>
                </button>
            </div>
        @endif

        @if (Session::has('message'))
            <div class="alert alert-info alert-dismissible fade show" role="alert">
                <strong>Message: </strong> {{ Session::get('message') }}
                <button type="button" class="close" data-dismiss="alert" aria-label="Close">
                    <span aria-hidden="true">&times;</span>
                </button>
            </div>
        @endif
        @if (Session::has('error_message'))
            <div class="alert alert-info alert-dismissible fade show" role="alert">
                <strong>Message: </strong> {{ Session::get('error_message') }}
                <button type="button" class="close" data-dismiss="alert" aria-label="Close">
                    <span aria-hidden="true">&times;</span>
                </button>
            </div>
        @endif
        <table border="2" class="table table-striped" id="table">
            <thead>
                <tr>
                    <th scope="col">No</th>
                    <th scope="col">Product Name</th>
                    <th scope="col">Product Description</th>
                    <th scope="col">Product Price</th>
                    <th scope="col">Product Image</th>
                    <th scope="col">Action</th>
                </tr>
            </thead>
            <tbody>
                @if (isset($products['message']) && is_array($products['message']) && count($products['message']) > 0)
                    @forelse ($products['message'] as $key => $item)
                        <tr>
                            <th>{{ $key + 1 }}</th>
                            <td>{{ $item['name'] }}</td>
                            <td>{{ $item['description'] }}</td>
                            <td>{{ $item['price'] }}</td>
                            <td><img src="http://127.0.0.1:8003/product/image/{{ $item['image'] }}" alt="Product Image"
                                    style="width: 30%; height: auto;">
                            </td>
                            <td>
                                <form id="" action="" method="POST">
                                    @csrf
                                    @method('DELETE')
                                    <a href="" class="btn btn-primary btn-sm"><i class="fas fa-eye"></i></a>
                                    <a href="/edit" class="btn btn-warning btn-sm mr-3 ml-3"><i
                                            class="fas fa-edit"></i></a>
                                    <button type="button" class="btn btn-danger btn-sm delete" name=""
                                        id="" onclick=""><i class="fas fa-trash"></i></button>
                                </form>
                            </td>
                        </tr>
                    @empty
                        <tr>
                            <td colspan="4">No
                                data available
                            </td>
                        </tr>
                    @endforelse
                @else
                    <tr>
                        <td colspan="4">No data available</td>
                    </tr>
                @endif
            </tbody>
        </table>
    </div>
@endsection
