@extends('admin.layout.welcome')

@section('title')
    Create Table
@endsection

@push('js')
    <script src="https://cdn.jsdelivr.net/npm/sweetalert2@11"></script>
    <script>
        function confirmCreate(event) {
            event.preventDefault();
            Swal.fire({
                title: 'Are you sure?',
                text: 'You will create this table!',
                icon: 'warning',
                showCancelButton: true,
                confirmButtonColor: '#d33',
                cancelButtonColor: '#3085d6',
                confirmButtonText: 'Yes, create it!'
            }).then((result) => {
                if (result.isConfirmed) {
                    document.getElementById('create-form').submit();
                }
            });
        }

        function validateForm(event) {
            confirmCreate(event);
        }
    </script>
@endpush

@section('content')
    <div class="container">
        @if (Session::has('error_message'))
            <div class="alert alert-danger alert-dismissible fade show" role="alert">
                <strong>Error: </strong> {{ Session::get('error_message') }}
                <button type="button" class="close" data-dismiss="alert" aria-label="Close">
                    <span aria-hidden="true">&times;</span>
                </button>
            </div>
        @endif
        <form action="{{ url('/admin/table') }}" method="POST" enctype="multipart/form-data" id="create-form">
            @csrf
            <div class="mb-3">
                <label for="number" class="form-label">Table Number <span style="color: red">*</span></label>
                <input type="text" class="form-control" id="number" name="number" value="{{ old('number') }}">
                @error('number')
                    <div class="alert alert-danger mt-1">{{ $message }}</div>
                @enderror
            </div>
            <div class="mb-3">
                <label for="capacity" class="form-label">Table Capacity <span style="color: red">*</span></label>
                <input type="number" class="form-control" id="capacity" name="capacity" value="{{ old('capacity') }}">
                @error('capacity')
                    <div class="alert alert-danger mt-1">{{ $message }}</div>
                @enderror
            </div>
            <button type="button" class="btn btn-primary" onclick="validateForm(event)">Create</button>
        </form>
    </div>
@endsection
