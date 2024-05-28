<?php

use Illuminate\Support\Facades\Route;
use App\Http\Controllers\Admin\AdminController;
use App\Http\Controllers\Admin\CategoryController;
use App\Http\Controllers\Admin\OrderController;
use App\Http\Controllers\Admin\ProductController;
use App\Http\Controllers\Admin\TableController;


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

Route::prefix('admin')->namespace('App\Http\Controllers\Admin')->group(function () {
    // Admin Login Route
    Route::match(['get', 'post'], 'login', [AdminController::class, 'adminLogin']);

    Route::middleware('auth.admin')->group(function () {
        // Admin Profile Route
        Route::get('/profile', [AdminController::class, 'getProfile'])->name('profile.admin');

        // Admin Logout Route
        Route::post('/logout', [AdminController::class, 'logoutAdmin'])->name('logout.submit');

        // Admin Dashboard Route
        Route::get('/dashboard', 'AdminController@Dashboard')->name('dashboard');

        // Manage Category
        Route::resource('category', CategoryController::class)->except(['update']);
        Route::put('category/{category}', [CategoryController::class, 'update'])->name('admin.category.update');


        // Manage Products
        Route::resource('product', ProductController::class);

        //Manage Table
        Route::resource('table', TableController::class)->except(['update']);
        Route::put('table/{table}', [TableController::class, 'update'])->name('admin.table.update');
        Route::put('order/approve{id}', 'AdminController@approve')->name('admin.approve');
        Route::put('order/reject{id}', 'AdminController@reject')->name('admin.reject');
    });

        // Manage Order
        Route::resource('order', OrderController::class);
});
