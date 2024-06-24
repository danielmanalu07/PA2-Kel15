@extends('admin.layout.welcome')
@section('title')
List Request Tables
@endsection

@push('js')
<script src="https://cdn.jsdelivr.net/npm/sweetalert2@11"></script>
<script>
    function confirmUpdateStatus(event, id, action) {
        event.preventDefault();
        Swal.fire({
            title: 'Are you sure?',
            text: 'Do you want to update the status of this request table?',
            icon: 'warning',
            showCancelButton: true,
            confirmButtonColor: '#3085d6',
            cancelButtonColor: '#d33',
            confirmButtonText: 'Yes, update it!'
        }).then((result) => {
            if (result.isConfirmed) {
                document.getElementById('update-status-form-' + id + '-' + action).submit();
            }
        });
    }

    $(document).ready(function() {
        $('#search').keyup(function() {
            var value = $(this).val().toLowerCase();
            $('#requestTable tbody tr').filter(function() {
                $(this).toggle($(this).text().toLowerCase().indexOf(value) > -1)
            });
        });
    });
</script>
@endpush

@section('content')
<div class="container">
    @if (Session::has('success'))
    <div class="alert alert-success alert-dismissible fade show" role="alert">
        <strong>Success: </strong> {{ Session::get('success') }}
        <button type="button" class="close" data-dismiss="alert" aria-label="Close">
            <span aria-hidden="true">&times;</span>
        </button>
    </div>
    @endif

    @if (Session::has('error'))
    <div class="alert alert-danger alert-dismissible fade show" role="alert">
        <strong>Error: </strong> {{ Session::get('error') }}
        <button type="button" class="close" data-dismiss="alert" aria-label="Close">
            <span aria-hidden="true">&times;</span>
        </button>
    </div>
    @endif

    @if (!empty($requestTables['message']) && is_array($requestTables['message']))
    <table class="table table-striped" id="requestTable">
        <thead>
            <tr>
                <th scope="col">ID</th>
                <th scope="col">Table ID</th>
                <th scope="col">Customer ID</th>
                <th scope="col">Status</th>
                <th scope="col">Notes</th>
                <th scope="col">Action</th>
            </tr>
        </thead>
        <tbody>
            @foreach($requestTables['message'] as $requestTable)
                <tr>
                    <td>{{ $requestTable['id'] }}</td>
                    <td>{{ $requestTable['table_id'] }}</td>
                    <td>{{ $requestTable['customer_id'] }}</td>
                    <td>{{ $requestTable['status'] == 1 ? 'Approved' : ($requestTable['status'] == 0 ? 'Pending' : 'Rejected') }}</td>
                    <td>{{ $requestTable['notes'] }}</td>
                    <td>
                        @if($requestTable['status'] == 0) <!-- Only show buttons if status is Pending -->
                            <form id="update-status-form-{{ $requestTable['id'] }}-approve" action="{{ route('admin.requestTable.approve', $requestTable['id']) }}" method="POST" style="display:inline;">
                                @csrf
                                @method('PUT')
                                <input type="hidden" name="status" value="1">
                                <button type="button" class="btn btn-success btn-sm" onclick="confirmUpdateStatus(event, '{{ $requestTable['id'] }}', 'approve')">
                                    Approve
                                </button>
                            </form>
                            <form id="update-status-form-{{ $requestTable['id'] }}-reject" action="{{ route('admin.requestTable.reject', $requestTable['id']) }}" method="POST" style="display:inline;">
                                @csrf
                                @method('PUT')
                                <input type="hidden" name="status" value="-1">
                                <button type="button" class="btn btn-danger btn-sm" onclick="confirmUpdateStatus(event, '{{ $requestTable['id'] }}', 'reject')">
                                    Reject
                                </button>
                            </form>
                        @endif
                    </td>
                </tr>
            @endforeach
        </tbody>
    </table>
    @else
    <p>No request tables found.</p>
    @endif
</div>
@endsection
