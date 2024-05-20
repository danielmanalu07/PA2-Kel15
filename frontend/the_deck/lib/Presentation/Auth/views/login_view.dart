import 'package:the_deck/Core/Routes/routes_name.dart';
import 'package:the_deck/Core/app_colors.dart';
import 'package:the_deck/Core/response_conf.dart';
import 'package:the_deck/Core/text_styles.dart';
import 'package:the_deck/Presentation/Auth/screens/default_button.dart';
import 'package:flutter/gestures.dart';
import 'package:flutter/material.dart';
import '../screens/default_field.dart';
import 'package:gap/gap.dart';
import 'package:flutter_svg/flutter_svg.dart';
import 'package:the_deck/Core/Utils/utils.dart';
class LoginView extends StatelessWidget {
  const LoginView({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    final emailController = TextEditingController();
    final passwordController = TextEditingController();

    MathUtils.init(context);
    return Scaffold(
        body: Padding(
      padding: EdgeInsets.symmetric(horizontal: getWidth(24)).copyWith(
        top: MediaQuery.of(context).viewPadding.top,
      ),
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          const Gap(32),
          Text(
            "Login to your\n account.",
            style: TextStyles.headingH4SemiBold
                .copyWith(color: Pallete.neutral100),
          ),
          const Gap(8),
          Text(
            "Please sign in to your account ",
            style:
                TextStyles.bodyMediumMedium.copyWith(color: Pallete.neutral60, fontSize: getFontSize(14)),
          ),
          const Gap(32),
          DefaultField(
            hintText: "Enter Username",
            controller: emailController,
            labelText: "Username",
          ),
          const Gap(14),
          DefaultField(
            hintText: "Password",
            controller: passwordController,
            labelText: "Password",
            isPasswordField: true,
          ),
          const Gap(24),
          Align(
            alignment: Alignment.topRight,
            child: InkWell(
              onTap: () => Navigator.pushNamed(context, RoutesName.forgetPassword),
              child: Text(
                "Forgot password?",
                style: TextStyles.bodyMediumMedium
                    .copyWith(color: Pallete.orangePrimary, fontSize: getFontSize(14)),
              ),
            ),
          ),
          const Gap(24),
          DefaultButton(
            btnContent: "Sign",
            function: () => Navigator.pushReplacementNamed(context, RoutesName.main),
          ),
          const Gap(24),
      
          const Gap(32),
          Align(
            alignment: Alignment.center,
            child: Text.rich(
              TextSpan(
                children: [
                  TextSpan(
                      text: "Don't have an account?",
                      style: TextStyles.bodyMediumMedium
                          .copyWith(color: Pallete.neutral100, fontSize: getFontSize(14)), ),
                   TextSpan(
                      text: ' ', style: TextStyles.bodyMediumSemiBold.copyWith(
                       fontSize: getFontSize(14)
                   )),
                  TextSpan(
                      recognizer: TapGestureRecognizer()..onTap
                      =()=>Navigator.pushReplacementNamed(context, RoutesName.signUp),
                      text: 'Register',
                      style: TextStyles.bodyMediumSemiBold
                          .copyWith(color: Pallete.orangePrimary, fontSize: getFontSize(14))),
                ],
              ),
            ),
          )
        ],
      ),
    ));
  }
}
