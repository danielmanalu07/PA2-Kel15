import 'dart:convert';
import 'package:http/http.dart' as http;
import 'package:user/Config/ApiConfig.dart';

class AuthController {
  Future<String?> register(String username, String password, String email) async {
    try {
      final response = await http.post(
        Uri.parse('${baseUrl}/api/register'),
        headers: <String, String>{
          'Content-Type': 'application/json; charset=UTF-8',
        },
        body: jsonEncode(<String, String>{
          'username': username,
          'password': password,
          'email': email,
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
