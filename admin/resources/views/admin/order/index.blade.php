@extends('admin.layout.welcome')
@section('title')
List Order
@endsection

@push('css')
<style>
    .status-pending {
        color: orange;
    }
    .status-approved {
        color: green;
    }
    .status-rejected {
        color: red;
    }
    .status-ready {
        color: blue;
    }
    .status-completed {
        color: grey;
    }
</style>
@endpush

@push('js')
<script src="https://cdn.jsdelivr.net/npm/sweetalert2@11"></script>
<script>
    function submitActionForm(event, orderId, action) {
        event.preventDefault();
        document.getElementById('action-' + orderId).value = action;
        document.getElementById('action-form-' + orderId).submit();
    }
</script>
@endpush

@section('content')
<div class="container-fluid pt-4 px-4">
    <div class="bg-light rounded h-100 p-4">
        <div class="table-responsive">
            <table class="table">
                <thead>
                    <tr>
                        <th scope="col">Code</th>
                        <th scope="col">Customer ID</th>
                        <th scope="col">Total</th>
                        <th scope="col">Note</th>
                        <th scope="col">Payment Method</th>
                        {{-- <th scope="col">Table ID</th> --}}
                        <th scope="col">Pick Up Type</th>
                        <th scope="col">Proof Of Payment</th>
                        <th scope="col">Status</th>
                        <th scope="col">Action</th>
                    </tr>
                </thead>
                <tbody>
                    @foreach ($orders as $order)
                    <tr>
                        <td>{{ $order['code'] }}</td>
                        <td>{{ $order['customer_id'] }}</td>
                        <td>{{ $order['total'] }}</td>
                        <td>{{ $order['note'] }}</td>
                        <td>{{ $order['payment_method'] }}</td>
                        {{-- <td>{{ $order['table_id'] }}</td> --}}
                        <td>{{ $order['pick_up_type'] }}</td>
                        <td>
                            @if($order['proof_of_payment'])
                                <img src="http://172.27.1.162:8080/order/image/{{ $order['proof_of_payment'] }}"
                                    alt="Proof of Payment" class="img-fluid rounded mb-3"
                                    style="max-width: 30%; height: auto;">
                            @else
                                <p>Bukti Pembayaran tidak tersedia</p>
                            @endif
                        </td>
                        <td class="{{ $order['status'] == 0 ? 'status-pending' : ($order['status'] == 1 ? 'status-approved' : ($order['status'] == 2 ? 'status-rejected' : ($order['status'] == 3 ? 'status-ready' : 'status-completed'))) }}">
                            @if($order['status'] == 0)
                            Waiting
                            @elseif($order['status'] == 1)
                            Accepted
                            @elseif($order['status'] == 2)
                            Rejected
                            @elseif($order['status'] == 3)
                            Ready for Pickup
                            @elseif($order['status'] == 4)
                            Finished
                            @else
                            Cancelled
                            @endif
                        </td>
                        <td>
                            @if(!in_array($order['status'], [2, 4]))
                            <form id="approve-form-{{ $order['id'] }}" action="{{ route('admin.approve', $order['id']) }}" method="POST" class="d-inline">
                                @csrf
                                @method('PUT')
                                <button class="btn btn-sm btn-primary">Approve</button>
                            </form>
                            @if($order['status'] != 1)
                                <form id="reject-form-{{ $order['id'] }}" action="{{ route('admin.reject', $order['id']) }}" method="POST" class="d-inline">
                                    @csrf
                                    @method('PUT')
                                    <button class="btn btn-sm btn-danger">Reject</button>
                                </form>
                            @endif
                            @if($order['status'] == 1)
                                <form id="ready-form-{{ $order['id'] }}" action="{{ route('admin.ready', $order['id']) }}" method="POST" class="d-inline">
                                    @csrf
                                    @method('PUT')
                                    <button class="btn btn-sm btn-info">Ready for Pickup</button>
                                </form>
                            @endif
                            @if($order['status'] == 3)
                                <form id="complete-form-{{ $order['id'] }}" action="{{ route('admin.complete', $order['id']) }}" method="POST" class="d-inline">
                                    @csrf
                                    @method('PUT')
                                    <button class="btn btn-sm btn-secondary">Complete</button>
                                </form>
                            @endif
                            @endif
                        </td>
                    </tr>
                    @endforeach
                </tbody>
            </table>
        </div>
    </div>
</div>
@endsection