import 'package:get/get.dart';
import 'package:the_deck/Core/Routes/routes_name.dart';
import 'package:the_deck/Guest/HomeView.dart';

import 'package:the_deck/Presentation/Auth/views/login_view.dart';
import 'package:the_deck/Presentation/Auth/views/signup_view.dart';
import 'package:the_deck/Presentation/Foods/Views/about_menu_view.dart';
import 'package:the_deck/Presentation/Main/main_view.dart';
import 'package:the_deck/Presentation/Profil/edit_personal_data_view.dart';
import 'package:the_deck/Presentation/Profil/personal_data_view.dart';
import 'package:flutter/material.dart';

class Routes {
  static Route<dynamic> onGenerateRoute(RouteSettings routeSettings) {
    switch (routeSettings.name) {
      case RoutesName.login:
        return MaterialPageRoute(builder: (context) => const LoginView());
      case RoutesName.signUp:
        return MaterialPageRoute(builder: (context) => const SignUpView());
      case RoutesName.main:
        return MaterialPageRoute(builder: (context) => const MainView());
      case RoutesName.guest:
        return MaterialPageRoute(builder: (context) => HomeGuest());

      // case RoutesName.aboutMenu:
      //   return MaterialPageRoute(builder: (context) => const AboutMenuView());

      case RoutesName.personnalData:
        return MaterialPageRoute(
            builder: (context) => const PersonalDataView());
      case RoutesName.EditPersonalDataView:
        return MaterialPageRoute(
            builder: (context) => const EditPersonalDataView());
      default:
        return MaterialPageRoute(
          builder: (context) => const Scaffold(
            body: Text("No routes found"),
          ),
        );
    }
  }
}
