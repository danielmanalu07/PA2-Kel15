import 'package:get/get.dart';
import 'package:the_deck/Core/Routes/routes.dart';
import 'package:the_deck/Core/Routes/routes_name.dart';
import 'package:flutter/material.dart';

void main() {
  runApp(const App());
}

class App extends StatelessWidget {
  const App({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return GetMaterialApp(
      debugShowCheckedModeBanner: false,
      initialRoute: RoutesName.login,
      onGenerateRoute: Routes.onGenerateRoute,
      theme: ThemeData(canvasColor: Colors.white),
    );
  }
}
