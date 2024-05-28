import 'dart:convert';
import 'package:flutter/material.dart';
import 'package:get/get.dart';
import 'package:http/http.dart' as http;
import 'package:the_deck/Models/Product.dart';

class ProductController extends GetxController {
  final product = Rxn<Product>();
  var productList = <Product>[].obs;

  @override
  void onInit() {
    super.onInit();
    fetchProductList();
  }

  Future<void> fetchProductList() async {
    final response =
        await http.get(Uri.parse('http://192.168.30.215:8080/product'));
    if (response.statusCode == 200) {
      final List<dynamic> data = json.decode(response.body)['message'];
      productList.value = data.map((json) => Product.fromJson(json)).toList();
    } else {
      throw Exception('Failed to load products');
    }
  }

  Future<List<Product>> getProductList() async {
    await fetchProductList();
    return productList.toList();
  }

  Future<void> getProductById(int id) async {
    final response =
        await http.get(Uri.parse('http://192.168.30.215:8080/product/${id}'));

    if (response.statusCode == 200) {
      final data = jsonDecode(response.body);
      product.value = Product.fromJson(data['message']);
    } else {
      Get.snackbar(
        'Error',
        'Could not fetch product data',
        snackPosition: SnackPosition.TOP,
        backgroundColor: Colors.red,
        colorText: Colors.white,
      );
      print('Fetching product data failed: ${response.body}');
    }
  }

  Future<void> getProductByCategoryId(int categoryId) async {
    final response = await http.get(
        Uri.parse('http://192.168.30.215:8080/product/category/$categoryId'));
    if (response.statusCode == 200) {
      final List<dynamic> data = json.decode(response.body)['message'] ?? [];
      productList.value = data.map((json) => Product.fromJson(json)).toList();
    } else {
      Get.snackbar(
        'Error',
        'Failed to load products for the category',
        snackPosition: SnackPosition.TOP,
        backgroundColor: Colors.red,
        colorText: Colors.white,
      );
      throw Exception('Failed to load products');
    }
  }
}
