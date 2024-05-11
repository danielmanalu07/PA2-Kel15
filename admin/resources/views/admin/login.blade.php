<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <link href='https://unpkg.com/boxicons@2.1.4/css/boxicons.min.css' rel='stylesheet'>
    <link rel="stylesheet" href="style.css">
    <title>Login Admin</title>
    <link rel="stylesheet" href="{{ asset('css/style.css') }}">
    <link rel="stylesheet" href="{{ asset('js/script.js') }}">
    <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.0.2/dist/css/bootstrap.min.css" rel="stylesheet"
        integrity="sha384-EVSTQN3/azprG1Anm3QDgpJLIm9Nao0Yz1ztcQTwFspd3yD65VohhpuuCOmLASjC" crossorigin="anonymous">
    <script src="https://cdn.jsdelivr.net/npm/bootstrap@5.0.2/dist/js/bootstrap.bundle.min.js"
        integrity="sha384-MrcW6ZMFYlzcLA8Nl+NtUVF0sA7MsXsP1UyJoMp4YLEuNSfAP+JcXn/tWtIaxVXM" crossorigin="anonymous">
    </script>
</head>

<body background="{{ asset('images/1.jpg') }}">
    <div class="wrapper">

        <div class="form-box">

            <!-- login form -->
            <div class="login-container" id="login">
                <div class="top">
                    <!-- Tambahkan kelas 'logo' pada img -->
                    <center>
                        @if (Session::has('error'))
                            <div class="alert alert-danger">
                                <strong>Error: </strong>{{ Session::get('error') }}
                            </div>
                        @endif
                        @if (Session::has('success'))
                            <div class="alert alert-success">
                                <strong>Success: </strong>{{ Session::get('success') }}
                            </div>
                        @endif
                        <img class="logo" src="{{ asset('images/Logo.jpeg') }}" alt="">
                    </center>
                </div>
                <form action="{{ url('/admin/login') }}" method="post">
                    @csrf
                    <div class="input-box">
                        <input type="text" class="input-field" placeholder="Username" name="username">
                        <i class="bx bx-user"></i>
                    </div>
                    @error('username')
                        <div class="alert alert-danger">{{ $message }}</div>
                    @enderror
                    <div class="input-box">
                        <input type="password" class="input-field" placeholder="Password" name="password">
                        <i class="bx bx-lock-alt"></i>
                    </div>
                    @error('password')
                        <div class="alert alert-danger">{{ $message }}</div>
                    @enderror
                    <div class="input-box">
                        <input type="submit" class="submit" value="Sign In">
                    </div>
                </form>
            </div>
        </div>
        <div class="animation"></div>
    </div>
</body>

</html>
