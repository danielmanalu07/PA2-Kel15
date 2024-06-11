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
@endpush

@section('content')
<div class="container">
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

    @if (!empty($tables['message']) && is_array($tables['message']))
    <table border="2" class="table table-striped" id="table">
        <thead>
            <tr>
                <th scope="col">ID</th>
                <th scope="col">Number</th>
                <th scope="col">Capacity</th>
                <th scope="col">Status</th>
                <th scope="col">Action</th>
            </tr>
        </thead>
        <tbody>
            @foreach($tables['message'] as $table)
                <tr>
                    <td>{{ $table['id'] }}</td>
                    <td>{{ $table['number'] }}</td>
                    <td>{{ $table['capacity'] }}</td>
                    <td>
                        {{ $table['status'] == 1 ? 'Sedang digunakan' : ($table['status'] == 0 ? 'Tersedia' : ($table['status'] == 2 ? 'Rejected' : 'Finished')) }}
                    </td>
                    <td>
                        <form id="delete-form-{{ $table['id'] }}" action="/admin/table/{{ $table['id'] }}" method="POST">
                            @csrf
                            @method('DELETE')
                            <a href="/admin/table/{{ $table['id'] }}" class="btn btn-primary btn-sm"><i class="fas fa-eye"></i> Show</a>
                            <a href="/admin/table/{{ $table['id'] }}/edit" class="btn btn-warning btn-sm mr-3 ml-3"><i class="fas fa-edit"></i> Edit</a>
                            <button type="button" class="btn btn-danger btn-sm delete" onclick="confirmDelete(event, '{{ $table['id'] }}')"><i class="fas fa-trash"></i> Delete</button>
                        </form>
                    </td>
                </tr>
            @endforeach
        </tbody>
    </table>
    @else
    <p>No tables found.</p>
    @endif
</div>
@endsection
