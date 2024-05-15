import 'package:flutter/material.dart';
import 'package:get/get.dart';
import 'package:user/Screens/Login.dart';
import 'package:user/Screens/admin/login.dart';

void main() {
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  @override
  Widget build(BuildContext context) {
    return GetMaterialApp(
      title: 'Flutter Demo',
      
      home: const LoginScreen(),
    );
  }
}
