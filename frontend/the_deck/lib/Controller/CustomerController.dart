import 'dart:convert';
import 'package:flutter/material.dart';
import 'package:get/get.dart';
import 'package:get_storage/get_storage.dart';
import 'package:http/http.dart' as http;
import 'package:the_deck/Models/Cart_Item.dart';
import 'package:the_deck/Models/Customer.dart';
import 'package:the_deck/Models/Product.dart';
import 'package:the_deck/Models/Register.dart';
import 'package:the_deck/Presentation/Auth/views/login_view.dart';
import 'package:the_deck/Presentation/Cart/cart_view.dart';
import 'package:the_deck/Presentation/Main/main_view.dart';

class RegisterController extends GetxController {
  final token = ''.obs;
  final box = GetStorage();
  final userProfile = Rxn<Customer>();
  var cartItems = <CartItem>[].obs;

  Future<void> registerUser(RegisterModel registerModel) async {
    final url = Uri.parse('http://192.168.30.215:8080/customer/register');

    var request = http.MultipartRequest('POST', url);
    request.headers.addAll({'Content-Type': 'multipart/form-data'});
    request.fields['name'] = registerModel.name;
    request.fields['username'] = registerModel.username;
    request.fields['email'] = registerModel.email;
    request.fields['password'] = registerModel.password;
    request.fields['phone'] = registerModel.phone;
    request.fields['address'] = registerModel.address;
    request.fields['gender'] = registerModel.gender;
    request.fields['DateOfBirth'] = registerModel.dateOfBirth;

    if (registerModel.image.isNotEmpty) {
      request.files
          .add(await http.MultipartFile.fromPath('image', registerModel.image));
    }

    final response = await request.send();
    final responseBody = await response.stream.bytesToString();

    if (response.statusCode == 200) {
      Get.offAll(() => LoginView());
      Get.snackbar(
        'Success',
        'Registration Successfully',
        snackPosition: SnackPosition.TOP,
        backgroundColor: Colors.green,
        colorText: Colors.white,
      );
      print('Registration successful');
    } else {
      Get.snackbar(
        'Error',
        'Could not Register',
        snackPosition: SnackPosition.TOP,
        backgroundColor: Colors.red,
        colorText: Colors.white,
      );
      print('Registration failed: $responseBody');
    }
  }

  Future<void> loginUser(String email, String password) async {
    final url = Uri.parse('http://192.168.30.215:8080/customer/login');
    final response = await http.post(
      url,
      headers: {'Content-Type': 'application/json'},
      body: jsonEncode({'Email': email, 'Password': password}),
    );

    if (response.statusCode == 200) {
      final responseData = jsonDecode(response.body);
      final token = responseData['token'];
      box.write('token', token);

      Get.offAll(() => MainView());
      Get.snackbar(
        'Success',
        'Login Successfully',
        snackPosition: SnackPosition.TOP,
        backgroundColor: Colors.green,
        colorText: Colors.white,
      );
      print('Login successful');
    } else {
      final responseBody = jsonDecode(response.body);
      Get.snackbar(
        'Error',
        'Could not Login',
        snackPosition: SnackPosition.TOP,
        backgroundColor: Colors.red,
        colorText: Colors.white,
      );
      print('Login failed: $responseBody');
    }
  }

  Future<void> getUserProfile() async {
    final url = Uri.parse('http://192.168.30.215:8080/customer/profile');
    final token = box.read('token');
    final response = await http.get(
      url,
      headers: {'Cookie': 'jwt=$token'},
    );

    print('Token : $token');

    if (response.statusCode == 200) {
      final responseData = jsonDecode(response.body);
      userProfile.value = Customer.fromJson(responseData['message']);
    } else {
      Get.snackbar(
        'Error',
        'Could not fetch profile data',
        snackPosition: SnackPosition.TOP,
        backgroundColor: Colors.red,
        colorText: Colors.white,
      );
      print('Fetching profile data failed: ${response.body}');
    }
  }

  Future<void> logout() async {
    final url = Uri.parse('http://192.168.30.215:8080/customer/logout');
    final token = box.read('token');
    final response = await http.post(
      url,
      headers: {'Cookie': 'jwt=$token'},
    );

    if (response.statusCode == 200) {
      Get.offAll(() => LoginView());
      Get.snackbar(
        'Success',
        'Logout Successfully',
        snackPosition: SnackPosition.TOP,
        backgroundColor: Colors.green,
        colorText: Colors.white,
      );
      print('Logout successful');
    } else {
      Get.snackbar(
        'Error',
        'Logout Error',
        snackPosition: SnackPosition.TOP,
        backgroundColor: Colors.red,
        colorText: Colors.white,
      );
      print('Logout Error');
    }
  }

  Future<void> addToCart(int productId, int quantity) async {
    final url = Uri.parse('http://192.168.30.215:8080/cart/add');
    final token = box.read('token');
    final response = await http.post(
      url,
      headers: {'Cookie': 'jwt=$token'},
      body: {
        'product_id': productId.toString(),
        'quantity': quantity.toString()
      },
    );

    if (response.statusCode == 200) {
      Get.to(() => CartView());
      Get.snackbar(
        'Success',
        'Add To Cart Successfully',
        snackPosition: SnackPosition.TOP,
        backgroundColor: Colors.green,
        colorText: Colors.white,
      );
      print('Add To Cart successful');
    } else {
      Get.snackbar(
        'Error',
        'Add To Cart Error',
        snackPosition: SnackPosition.TOP,
        backgroundColor: Colors.red,
        colorText: Colors.white,
      );
      print('Add To Cart Error');
    }
  }

  Future<void> getMyCart() async {
    final url = Uri.parse('http://192.168.30.215:8080/cart/myCart');
    final token = box.read('token');
    final response = await http.get(
      url,
      headers: {'Cookie': 'jwt=$token'},
    );

    if (response.statusCode == 200) {
      final responseData = jsonDecode(response.body);
      List<dynamic> data = responseData['message'];
      List<CartItem> items =
          data.map((json) => CartItem.fromJson(json)).toList();
      cartItems.assignAll(items);
    } else {
      print('Failed to fetch cart items');
    }
  }

  Future<void> deleteCartItem(int cartItemId) async {
    final url = Uri.parse('http://192.168.30.215:8080/cart/delete/$cartItemId');
    final token = box.read('token');
    final response = await http.delete(
      url,
      headers: {'Cookie': 'jwt=$token'},
    );

    if (response.statusCode == 200) {
      Get.snackbar(
        'Success',
        'Item deleted from cart',
        snackPosition: SnackPosition.TOP,
        backgroundColor: Colors.green,
        colorText: Colors.white,
      );
      cartItems.removeWhere((item) => item.id == cartItemId);
      print('Item deleted successfully');
    } else {
      Get.snackbar(
        'Error',
        'Failed to delete item from cart',
        snackPosition: SnackPosition.TOP,
        backgroundColor: Colors.red,
        colorText: Colors.white,
      );
      print('Failed to delete item: ${response.body}');
    }
  }

  Future<void> updateCartItemQuantity(int cartItemId, int quantity) async {
    final url = Uri.parse('http://192.168.30.215:8080/cart/edit/$cartItemId');
    final token = box.read('token');
    final response = await http.put(
      url,
      headers: {'Cookie': 'jwt=$token'},
      body: {'quantity': quantity.toString()},
    );

    if (response.statusCode == 200) {
      await getMyCart();
      print('Item quantity updated successfully');
    } else {
      print('Failed to update item quantity: ${response.body}');
    }
  }
}
