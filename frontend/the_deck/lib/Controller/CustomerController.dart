import 'dart:convert';
import 'dart:io';
import 'package:flutter/material.dart';
import 'package:get/get.dart';
import 'package:get_storage/get_storage.dart';
import 'package:http/http.dart' as http;
import 'package:image_picker/image_picker.dart';
import 'package:the_deck/Models/Cart_Item.dart';
import 'package:the_deck/Models/Customer.dart';
import 'package:the_deck/Models/Order.dart';
import 'package:the_deck/Models/Product.dart';
import 'package:the_deck/Models/Register.dart';
import 'package:the_deck/Models/RequestTable.dart';
import 'package:the_deck/Presentation/Auth/views/login_view.dart';
import 'package:the_deck/Presentation/Cart/MyOrder.dart';
import 'package:the_deck/Presentation/Cart/Order.dart';
import 'package:the_deck/Presentation/Cart/cart_view.dart';
import 'package:the_deck/Presentation/Main/home_view.dart';
import 'package:the_deck/Presentation/Main/main_view.dart';
import 'package:the_deck/Presentation/Profil/profil_view.dart';

class RegisterController extends GetxController {
  final token = ''.obs;
  final box = GetStorage();
  final userProfile = Rxn<Customer>();
  var cartItems = <CartItem>[].obs;
  var orderItems = <Order>[].obs;
  var reqTable = <RequestTable>[].obs;
  var isLoading = true.obs;

  Future<void> registerUser(RegisterModel registerModel) async {
    final url = Uri.parse('http://192.168.66.215:8080/customer/register');

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
    final url = Uri.parse('http://192.168.66.215:8080/customer/login');
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
    final url = Uri.parse('http://192.168.66.215:8080/customer/profile');
    final token = box.read('token');
    final response = await http.get(
      url,
      headers: {'Cookie': 'jwt=$token'},
    );

    if (response.statusCode == 200) {
      final responseData = jsonDecode(response.body);
      userProfile.value = Customer.fromJson(responseData['message'] ?? []);
    } else {
      // Get.snackbar(
      //   'Error',
      //   'Could not fetch profile data',
      //   snackPosition: SnackPosition.TOP,
      //   backgroundColor: Colors.red,
      //   colorText: Colors.white,
      // );
      print("Failed Data User");
    }
  }

  Future<void> logout() async {
    final url = Uri.parse('http://192.168.66.215:8080/customer/logout');
    final token = box.read('token');
    final response = await http.post(
      url,
      headers: {'Cookie': 'jwt=$token'},
    );

    if (response.statusCode == 200) {
      box.remove('token'); // Clear the token on logout
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
    final url = Uri.parse('http://192.168.66.215:8080/cart/add');
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
    } else if (response.statusCode == 401) {
      Get.snackbar(
        'Error',
        'You must login first',
        snackPosition: SnackPosition.TOP,
        backgroundColor: Colors.red,
        colorText: Colors.white,
      );
      Get.to(() => LoginView());
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

  Future<void> requestTable(int selectedTableId, String notes) async {
    final token = GetStorage().read('token');
    final url = 'http://192.168.66.215:8080/requestTable/create';
    final headers = {
      'Cookie': 'jwt=$token',
      'Content-Type': 'application/json',
    };
    final body = json.encode({
      'table_id': selectedTableId,
      'notes': notes,
    });

    final response =
        await http.post(Uri.parse(url), headers: headers, body: body);

    if (response.statusCode == 200) {
      print(response.body);
      Get.snackbar('Success', 'Table requested successfully');
      Get.to(() => ProfilView());
    } else if (response.statusCode == 401) {
      Get.snackbar(
        'Error',
        'You must login first',
        snackPosition: SnackPosition.TOP,
        backgroundColor: Colors.red,
        colorText: Colors.white,
      );
      Get.to(() => LoginView());
    } else {
      print(response.body);
      Get.snackbar('Error', 'Failed to request table');
    }
  }

  Future<void> getMyOrder() async {
    final url = Uri.parse('http://192.168.66.215:8080/order/myorder');
    final token = box.read('token');
    final response = await http.get(
      url,
      headers: {'Cookie': 'jwt=$token'},
    );

    if (response.statusCode == 200) {
      final responseData = jsonDecode(response.body);
      List<dynamic> data = responseData['message'] ?? [];
      List<Order> items = data.map((json) => Order.fromJson(json)).toList();
      orderItems.assignAll(items);
    } else {
      print('Failed to fetch order');
    }
  }

  Future<void> getMyRequestTable() async {
    final url = Uri.parse('http://192.168.66.215:8080/requestTable/myRequest');
    final token =
        box.read('token'); // Pastikan Anda telah menginisialisasi 'box'
    final response = await http.get(
      url,
      headers: {'Cookie': 'jwt=$token'},
    );

    if (response.statusCode == 200) {
      final responseData = jsonDecode(response.body);
      List<dynamic> data = responseData['message'] ?? [];
      List<RequestTable> items =
          data.map((json) => RequestTable.fromJson(json)).toList();
      reqTable.assignAll(
          items); // Ganti 'reqTable' menjadi 'customerController.reqTable'
    } else {
      print("Failed to fetch Request Table");
    }

    isLoading.value =
        false; // Ganti 'isLoading' menjadi 'customerController.isLoading'
  }

  Future<void> getMyCart() async {
    final url = Uri.parse('http://192.168.66.215:8080/cart/myCart');
    final token = box.read('token');
    final response = await http.get(
      url,
      headers: {'Cookie': 'jwt=$token'},
    );

    if (response.statusCode == 200) {
      final responseData = jsonDecode(response.body);
      List<dynamic> data = responseData['message'] ?? [];
      List<CartItem> items =
          data.map((json) => CartItem.fromJson(json)).toList();
      cartItems.assignAll(items);
    } else {
      print('Failed to fetch cart items');
    }
  }

  Future<void> deleteCartItem(int cartItemId) async {
    final url = Uri.parse('http://192.168.66.215:8080/cart/delete/$cartItemId');
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

  Future<void> updateCartItemQuantity(int cartItemId, int quantity,
      {VoidCallback? onSuccess}) async {
    final url = Uri.parse('http://192.168.66.215:8080/cart/edit/$cartItemId');
    final token = box.read('token');
    final response = await http.put(
      url,
      headers: {'Cookie': 'jwt=$token'},
      body: {'quantity': quantity.toString()},
    );

    if (response.statusCode == 200) {
      if (onSuccess != null) {
        onSuccess();
      }
      print('Item quantity updated successfully');
    } else {
      print('Failed to update item quantity: ${response.body}');
    }
  }

  Future<void> uploadImage(
      BuildContext context, int orderId, XFile? image) async {
    final token = box.read('token');
    if (image != null) {
      File imageFile = File(image.path);
      final uploadUrl = 'http://192.168.66.215:8080/order/payment/$orderId';

      var request = http.MultipartRequest('PUT', Uri.parse(uploadUrl));
      request.headers.addAll({'Cookie': 'jwt=${token}'});

      request.files
          .add(await http.MultipartFile.fromPath('image', imageFile.path));

      var response = await request.send();

      if (response.statusCode == 200) {
        print('Image uploaded successfully');
        ScaffoldMessenger.of(context).showSnackBar(
            SnackBar(content: Text('Image uploaded successfully')));
      } else {
        print('Image upload failed');
        ScaffoldMessenger.of(context)
            .showSnackBar(SnackBar(content: Text('Image upload failed')));
      }
    } else {
      print('No image selected');
      ScaffoldMessenger.of(context)
          .showSnackBar(SnackBar(content: Text('No image selected')));
    }
  }

  Future<void> updateOrderStatus(int orderId) async {
    try {
      final url = Uri.parse('http://192.168.66.215:8080/order/status/$orderId');
      final token = box.read('token');
      final response = await http.put(
        url,
        headers: {'Cookie': 'jwt=$token'},
        body: {'status': '5'},
      );

      if (response.statusCode == 200) {
        print('Order status updated successfully');
        Get.snackbar(
          'Success',
          'Canceled Successfully',
          snackPosition: SnackPosition.TOP,
          backgroundColor: Colors.green,
          colorText: Colors.white,
        );
      } else {
        print('Failed to update order status: ${response.body}');
      }
    } catch (e) {
      print('Error updating order status: $e');
    }
  }

  Future<void> CancleRequestTable(int reqTableId) async {
    try {
      final url = Uri.parse(
          'http://192.168.66.215:8080/requestTable/status/$reqTableId');
      final token = box.read('token');
      final response = await http.put(
        url,
        headers: {'Cookie': 'jwt=$token'},
        body: {'status': '2'},
      );

      if (response.statusCode == 200) {
        print('Request Table status updated successfully');
        Get.snackbar(
          'Success',
          'Canceled Successfully',
          snackPosition: SnackPosition.TOP,
          backgroundColor: Colors.green,
          colorText: Colors.white,
        );
        await getMyRequestTable();
      } else {
        print('Failed to update req. Table status: ${response.body}');
      }
    } catch (e) {
      print('Error updating req. Table  status: $e');
    }
  }

  Future<void> updateUserProfile(Customer updatedCustomer) async {
    final url = Uri.parse('http://192.168.66.215:8080/customer/update-profile');
    final token = box.read('token');
    final response = await http.put(
      url,
      headers: {
        'Content-Type': 'application/json',
        'Cookie': 'jwt=$token',
      },
      body: jsonEncode(updatedCustomer.toJson()),
    );

    if (response.statusCode == 200) {
      Get.offAll(() => MainView());
      Get.snackbar(
        'Success',
        'Profile updated successfully',
        snackPosition: SnackPosition.TOP,
        backgroundColor: Colors.green,
        colorText: Colors.white,
      );
    } else {
      Get.snackbar(
        'Error',
        'Could not update profile',
        snackPosition: SnackPosition.TOP,
        backgroundColor: Colors.red,
        colorText: Colors.white,
      );
    }
  }

  Future<String?> uploadImages(File image) async {
    final url = Uri.parse('http://192.168.66.215:8080/customer/upload-profile');
    final token = box.read('token');
    final request = http.MultipartRequest('POST', url);
    request.headers['Cookie'] = 'jwt=$token';

    request.files.add(await http.MultipartFile.fromPath('image', image.path));

    final response = await request.send();

    if (response.statusCode == 200) {
      final responseData = await http.Response.fromStream(response);
      final data = jsonDecode(responseData.body);
      return data['imageUrl'];
    } else {
      Get.snackbar(
        'Error',
        'Image upload failed',
        snackPosition: SnackPosition.TOP,
        backgroundColor: Colors.red,
        colorText: Colors.white,
      );
      return null;
    }
  }
}
