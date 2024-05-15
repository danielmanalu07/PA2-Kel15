import 'package:flutter/material.dart';
import 'package:the_deck/view/login_page/widgets/text_field_widget/text_field_input.dart';
import 'package:the_deck/view/login_page/widgets/text_field_widget/text_field_password.dart';

class RegisterTextField extends StatefulWidget {
  const RegisterTextField({Key? key}) : super(key: key);

  @override
  _RegisterTextFieldState createState() => _RegisterTextFieldState();
}

class _RegisterTextFieldState extends State<RegisterTextField> {
  @override
  Widget build(BuildContext context) {
    return Container(
        child: Column(
            mainAxisAlignment: MainAxisAlignment.center,
            crossAxisAlignment: CrossAxisAlignment.center,
            children: [
              TextFieldInput(text: "username", iconName: Icons.account_circle, ltext: "User Name"),
              TextFieldInput(text: "email", iconName: Icons.mail, ltext: "Email",),
              TextFieldPassword(),
              TextFieldPassword(),
            ]
        )
    );
  }
}
