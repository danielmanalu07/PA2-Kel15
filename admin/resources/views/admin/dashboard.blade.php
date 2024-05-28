@extends('admin.layout.welcome')
@section('title')
    Dashboard
@endsection
@push('js')
    @if (Session::has('success'))
        <script src="https://cdn.jsdelivr.net/npm/sweetalert2@10"></script>
        <script>
            document.addEventListener('DOMContentLoaded', function() {
                Swal.fire({
                    title: 'Selamat Datang!',
                    text: 'Selamat datang di dashboard admin.',
                    icon: 'success',
                    confirmButtonText: 'OK'
                });
            });
        </script>
    @endif
@endpush
@section('content')
    @php
        $currentTime = date('H');
        $greeting = '';
        if ($currentTime < 12.0) {
            $greeting = 'Selamat pagi';
        } elseif ($currentTime < 18.0 && $currentTime > 12.0) {
            $greeting = 'Selamat siang';
        } else {
            $greeting = 'Selamat malam';
        }
    @endphp

<div class="row-xl-3 col-md-12 mb-4">
    <h1 class="text-dark">{{ $greeting }}, {{ $data['message']['username'] }}!</h1>
</div> <br>
<div class="col-xl-4 col-md-6 mb-4">
    <a href="#" style="text-decoration: none; color: inherit;">
        <div class="card border-left-primary shadow h-100 py-2">
            <div class="card-body">
                <div class="row no-gutters align-items-center">
                    <div class="col mr-2">
                        <div class="text-xs font-weight-bold text-primary text-uppercase mb-1">
                            Order </div>
                        <div class="h5 mb-0 font-weight-bold text-gray-800">5</div>
                    </div>
                    <div class="col-auto">
                        <i class="fas fa-calendar fa-2x text-primary"></i>
                    </div>
                </div>
            </div>
        </div>
    </a>
</div>
<div class="col-xl-4 col-md-6 mb-4">
    <a href="product" style="text-decoration: none; color: inherit;">
        <div class="card border-left-danger shadow h-100 py-2">
            <div class="card-body">
                <div class="row no-gutters align-items-center">
                    <div class="col mr-2">
                        <div class="text-xs font-weight-bold text-danger text-uppercase mb-1">
                            Product </div>
                        <div class="h5 mb-0 font-weight-bold text-gray-800">Lihat Produk</div>
                    </div>
                    <div class="col-auto">
                        <i class="fas fa-clipboard-list fa-2x text-danger"></i>
                    </div>
                </div>
            </div>
        </div>
    </a>
</div>
<div class="col-xl-4 col-md-6 mb-4">
    <a href="#" style="text-decoration: none; color: inherit;">
        <div class="card border-left-success shadow h-100 py-2">
            <div class="card-body">
                <div class="row no-gutters align-items-center">
                    <div class="col mr-2">
                        <div class="text-xs font-weight-bold text-success text-uppercase mb-1">
                            Pendapatan</div>
                            <div class="h5 mb-0 font-weight-bold text-gray-800">$155135</div>
                    </div>
                    <div class="col-auto">
                        <i class="fas fa-dollar-sign fa-2x text-success"></i>
                    </div>
                </div>
            </div>
        </div>
    </a>
</div>

    
@endsection
