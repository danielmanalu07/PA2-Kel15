import 'package:flutter/material.dart';
import 'package:http/http.dart' as http;
import 'dart:convert';

import 'package:user/Config/ApiConfig.dart';

class AuthController {
  Future<String?> login(String username, String password) async {
    try {
      final response = await http.post(
        Uri.parse('${baseUrl}/customer/login'),
        headers: <String, String>{
          'Content-Type': 'application/json; charset=UTF-8',
        },
        body: jsonEncode(<String, String>{
          'username': username,
          'password': password,
        }),
      );

      if (response.statusCode == 200) {
        return "Login Successful";
      } else {
        final jsonResp = jsonDecode(response.body);
        final message = jsonResp['message'] as String?;
        return message;
      }
    } catch (e) {
      print('Error during login: $e');
      return null;
    }
  }

  Future<String?> register(String username, String password, String phone, String address) async {
    try {
      final response = await http.post(
        Uri.parse('${baseUrl}/customer/register'),
        headers: <String, String>{
          'Content-Type': 'application/json; charset=UTF-8',
        },
        body: jsonEncode(<String, String>{
          'username': username,
          'password': password,
          'phone': phone,
          'address': address,
        }),
      );

      if (response.statusCode == 200) {
        final jsonResp = jsonDecode(response.body);
        final message = jsonResp['message'] as String?;
        return message;
      } else {
        return null;
      }
    } catch (e) {
      print('Error during registration: $e');
      return null;
    }
  }
}
