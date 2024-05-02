import 'dart:convert';
import 'package:http/http.dart' as http;
import 'package:user/Config/ApiConfig.dart';

class AuthController {
  Future<String?> login(String username, String password) async {
    try {
      final response = await http.post(
        Uri.parse('${baseUrl}/login'),
        headers: <String, String>{
          'Content-Type': 'application/json; charset=UTF-8',
        },
        body: jsonEncode(<String, String>{
          'username': username,
          'password': password,
        }),
      );

      print('HTTP Response Code: ${response.statusCode}');
      print('HTTP Response Body: ${response.body}');

      if (response.statusCode == 200) {
        final jsonResp = jsonDecode(response.body);
        final token = jsonResp['token'] as String?;
        return token;
      } else {
        return null;
      }
    } catch (e) {
      print('Error during login: $e');
      return null;
    }
  }
}
