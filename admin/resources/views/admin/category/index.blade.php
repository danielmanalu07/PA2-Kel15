@extends('admin.layout.welcome')
@section('title')
    List Category
@endsection
@push('js')
    <script src="https://cdn.jsdelivr.net/npm/sweetalert2@11"></script>
    <script>
        function confirmDelete(event, id) {
            event.preventDefault();
            Swal.fire({
                title: 'Are you sure?',
                text: 'You will not be able to recover this data!',
                icon: 'warning',
                showCancelButton: true,
                confirmButtonColor: '#d33',
                cancelButtonColor: '#3085d6',
                confirmButtonText: 'Yes, delete it!'
            }).then((result) => {
                if (result.isConfirmed) {
                    document.getElementById('delete-form-' + id).submit();
                }
            });
        }

        $(document).ready(function() {
            $('#search').keyup(function() {
                var value = $(this).val().toLowerCase();
                $('#table tbody tr').filter(function() {
                    $(this).toggle($(this).text().toLowerCase().indexOf(value) > -1)
                });
            });
        });
    </script>
    <script></script>
@endpush
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
                    <th scope="col">Category Name</th>
                    <th scope="col">Category Description</th>
                    <th scope="col">Action</th>
                </tr>
            </thead>
            <tbody>
                @if (isset($category['message']) && is_array($category['message']) && count($category['message']) > 0)
                    @forelse ($category['message'] as $key => $item)
                        <tr>
                            <th>{{ $key + 1 }}</th>
                            <td>{{ $item['name'] }}</td>
                            <td>{{ $item['description'] }}</td>
                            <td>
                                <form id="delete-form-{{ $item['id'] }}" action="/admin/category/{{ $item['id'] }}"
                                    method="POST">
                                    @csrf
                                    @method('DELETE')
                                    <a href="/admin/category/{{ $item['id'] }}" class="btn btn-primary btn-sm"><i
                                            class="fas fa-eye"></i></a>
                                    <a href="/admin/category/{{ $item['id'] }}/edit"
                                        class="btn btn-warning btn-sm mr-3 ml-3"><i class="fas fa-edit"></i></a>
                                    <button type="button" class="btn btn-danger btn-sm delete" name=""
                                        id="" onclick="confirmDelete(event, '{{ $item['id'] }}')"><i
                                            class="fas fa-trash"></i></button>
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
