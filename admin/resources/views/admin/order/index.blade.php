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
                            <th scope="col">Table ID</th>
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
                                <td>{{ $order['table__id'] }}</td>
                                <td>{{ $order['pick_up_type'] }}</td>
                                <td><img src="http://192.168.187.215:8080/order/image/{{ $order['proof_of_payment'] }}"
                                        alt="Product Image" class="img-fluid rounded mb-3"
                                        style="max-width: 30%; height: auto;"></td>
                                <td>
                                    @if ($order['status'] == 0)
                                        Pending
                                    @elseif($order['status'] == 1)
                                        Approved
                                    @elseif($order['status'] == 2)
                                        Rejected
                                    @else
                                        Unknown
                                    @endif
                                </td>
                                <td>
                                    <form id="action-form-{{ $order['id'] }}"
                                        action="{{ route('admin.approve', $order['id']) }}" method="POST">
                                        @csrf
                                        @method('PUT')
                                        <button class="btn btn-sm btn-primary">Approve</button>
                                    </form>
                                    <form id="action-form-{{ $order['id'] }}"
                                        action="{{ route('admin.reject', $order['id']) }}" method="POST">
                                        @csrf
                                        @method('PUT')
                                        <button class="btn btn-sm btn-danger">Reject</button>
                                    </form>
                                </td>
                            </tr>
                        @endforeach
                    </tbody>
                </table>
            </div>
        </div>
    </div>
@endsection
