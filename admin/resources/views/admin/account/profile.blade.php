@extends('admin.layout.welcome')

@section('content')
    <div class="container mt-5">
        <div class="row justify-content-center">
            <div class="col-md-6">
                <div class="card">
                    <div class="card-header text-center">
                        <h3>Welcome, {{ $data['admin']['username'] }}</h3>
                        <span class="fas fa-user fa-3x"></span> <!-- Font Awesome user icon -->
                    </div>
                    <div class="card-body text-center">
                        <label for="">Username : </label>
                        <input type="text" readonly placeholder="{{ $data['admin']['username'] }}"> <br>
                    </div>
                    <div class="card-body">
                        <p class="text-center">Selamat datang di dashboard admin</p>
                        <p class="text-center">Silakan gunakan menu di sebelah kiri untuk mengakses fitur yang tersedia.</p>
                    </div>
                </div>
            </div>
        </div>
    </div>
@endsection
