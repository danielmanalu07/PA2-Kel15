@extends('admin.layout.welcome')
@section('title')
List Tables
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
    <div class="alert alert-danger alert-dismissible fade show" role="alert">
        <strong>Error: </strong> {{ Session::get('error_message') }}
        <button type="button" class="close" data-dismiss="alert" aria-label="Close">
            <span aria-hidden="true">&times;</span>
        </button>
    </div>
    @endif

    @if (!empty($tables) && isset($tables['message']) && is_array($tables['message']))
    <table border="2" class="table table-striped">
        <thead>
            <tr>
                <th scope="col">ID</th>
                <th scope="col">Number</th>
                <th scope="col">Capacity</th>
                <th scope="col">Actions</th>
            </tr>
        </thead>
        <tbody>
            @forelse($tables['message'] as $table)
                @if (isset($table['id']) && isset($table['number']) && isset($table['capacity']))
                    <tr>
                        <td>{{ $table['id'] }}</td>
                        <td>{{ $table['number'] }}</td>
                        <td>{{ $table['capacity'] }}</td>
                        <td>
                            <a href="" class="btn btn-primary btn-sm"><i class="fas fa-eye"></i></a>
                            <a href="{{ route('table.update', $table['id']) }}" class="btn btn-warning btn-sm mr-3 ml-3"><i
                                class="fas fa-edit"></i></i></a>
                            <form action="{{ route('table.destroy', $table['id']) }}" method="POST" style="display:inline-block;">
                                @csrf
                                @method('DELETE')
                                <button type="submit" class="btn btn-danger btn-sm" onclick="confirmDelete(event, '{{ $table['id'] }}')"><i class="fas fa-trash"></i></button>
                            </form>
                        </td>
                    </tr>
                @else
                    <tr>
                        <td colspan="4">Invalid table data</td>
                    </tr>
                @endif
            @empty
                <tr>
                    <td colspan="4">No tables found</td>
                </tr>
            @endforelse
        </tbody>
    </table>
    @else
        <p>No tables found.</p>
    @endif
</div>
@endsection
