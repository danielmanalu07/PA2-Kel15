import 'package:flutter/material.dart';
import 'package:the_deck/view/splash_screen.dart';
import 'package:the_deck/view/start_page/start_page_view.dart';

void main() {
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'The deck',
      debugShowCheckedModeBanner: false,
      theme: ThemeData(
        primarySwatch: Colors.orange,
      ),
      home: StartPageView(),
    );
  }
}
