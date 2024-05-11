<?php

use Illuminate\Support\Facades\Route;

/*
|--------------------------------------------------------------------------
| Web Routes
|--------------------------------------------------------------------------
|
| Here is where you can register web routes for your application. These
| routes are loaded by the RouteServiceProvider and all of them will
| be assigned to the "web" middleware group. Make something great!
|
 */

Route::get('/', function () {
    return redirect('/admin/login');
});

Route::prefix('/admin')->namespace('App\Http\Controllers\Admin')->group(function () {
    Route::match(['get', 'post'], 'login', 'AdminController@Login');
    Route::middleware('auth.admin')->group(function () {
        //Auth
        Route::get('/dashboard', 'AdminController@Dashboard');
        Route::get('/logout', 'AdminController@Logout')->name('logout.submit');
        Route::get('/profile', 'AdminController@Profile')->name('profile.admin');

        //Manage Category
        Route::resource('category', 'CategoryController');

        //Manage Products
        Route::resource('product', 'ProductController');
    });
});
