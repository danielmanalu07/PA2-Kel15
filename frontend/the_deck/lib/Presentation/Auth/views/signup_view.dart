// import 'dart:io';

// import 'package:the_deck/Core/Routes/routes_name.dart';
// import 'package:the_deck/Core/Utils/utils.dart';
// import 'package:the_deck/Core/app_colors.dart';
// import 'package:the_deck/Core/font_size.dart';
// import 'package:the_deck/Core/response_conf.dart';
// import 'package:the_deck/Core/text_styles.dart';
// import 'package:the_deck/Presentation/Auth/screens/account_status.dart';
// import 'package:the_deck/Presentation/Auth/screens/default_button.dart';
// import 'package:the_deck/Presentation/Auth/screens/default_field.dart';
// import 'package:flutter/gestures.dart';
// import 'package:flutter/material.dart';
// import 'package:flutter_svg/flutter_svg.dart';
// import 'package:gap/gap.dart';
// import 'package:image_picker/image_picker.dart';

// class SignUpView extends StatefulWidget {
//   const SignUpView({Key? key}) : super(key: key);

//   @override
//   _SignUpViewState createState() => _SignUpViewState();
// }

// class _SignUpViewState extends State<SignUpView> {
//   final ImagePicker _picker = ImagePicker();
//   XFile? _image;

//   @override
//   Widget build(BuildContext context) {
//     MathUtils.init(context);

//     return Scaffold(
//       resizeToAvoidBottomInset: false,
//       body: Padding(
//         padding: EdgeInsets.symmetric(horizontal: getWidth(24)).copyWith(
//           top: MediaQuery.of(context).viewPadding.top,
//         ),
//         child: SingleChildScrollView(
//           child: Column(
//             crossAxisAlignment: CrossAxisAlignment.start,
//             children: [
//               const Gap(32),
//               Text(
//                 "Create your new \naccount",
//                 style: TextStyles.headingH4SemiBold.copyWith(
//                   color: Pallete.neutral100,
//                   fontSize: getFontSize(FontSizes.h4),
//                 ),
//               ),
//               const Gap(8),
//               Text(
//                 "Create an account to start looking for the food \nyou like",
//                 style: TextStyles.bodyMediumMedium.copyWith(
//                   color: Pallete.neutral60,
//                   fontSize: getFontSize(FontSizes.medium),
//                 ),
//               ),
//               const Gap(12),
//               DefaultField(
//                 hintText: "Full Name",
//                 labelText: "Full Name",
//               ),
//               const Gap(14),
//               DefaultField(
//                 hintText: "Enter Email",
//                 labelText: "Email Address",
//               ),
//               const Gap(14),
//               DefaultField(
//                 hintText: "User Name",
//                 labelText: "User Name",
//               ),
//               const Gap(14),
//               DefaultField(
//                 hintText: "Password",
//                 labelText: "Password",
//                 isPasswordField: true,
//               ),
//               const Gap(14),
//               DefaultField(
//                 hintText: "Phone Number",
//                 labelText: "Phone Number",
//               ),
//               const Gap(14),
//               DefaultField(
//                 hintText: "Address",
//                 labelText: "Address",
//               ),
//               const Gap(14),
//               DefaultField(
//                 hintText: "Gender",
//                 labelText: "Gender",
//               ),
//               const Gap(14),
//               GestureDetector(
//                 onTap: () async {
//                   DateTime? pickedDate = await showDatePicker(
//                     context: context,
//                     initialDate: DateTime.now(),
//                     firstDate: DateTime(1900),
//                     lastDate: DateTime(2100),
//                   );
//                 },
//                 child: AbsorbPointer(
//                   child: DefaultField(
//                     hintText: "Date of Birth (YYYY-MM-DD)",
//                     labelText: "Date of Birth",
//                   ),
//                 ),
//               ),
//               const Gap(14),
//               GestureDetector(
//                 onTap: () async {
//                   final XFile? pickedFile = await _picker.pickImage(
//                     source: ImageSource.gallery,
//                   );
//                   setState(() {
//                     _image = pickedFile;
//                   });
//                 },
//                 child: Container(
//                   width: double.infinity,
//                   height: getHeight(150),
//                   decoration: BoxDecoration(
//                     color: Pallete.neutral20,
//                     borderRadius: BorderRadius.circular(8),
//                     border: Border.all(color: Pallete.neutral40),
//                   ),
//                   child: _image != null
//                       ? Image.file(
//                           File(_image!.path),
//                           fit: BoxFit.cover,
//                         )
//                       : Column(
//                           mainAxisAlignment: MainAxisAlignment.center,
//                           children: [
//                             Icon(
//                               Icons.camera_alt,
//                               color: Pallete.neutral60,
//                               size: getSize(50),
//                             ),
//                             const Gap(8),
//                             Text(
//                               "Upload Profile Picture",
//                               style: TextStyles.bodyMediumMedium.copyWith(
//                                 color: Pallete.neutral60,
//                                 fontSize: getFontSize(FontSizes.medium),
//                               ),
//                             ),
//                           ],
//                         ),
//                 ),
//               ),
//               const Gap(24),
//               const Gap(24),
//               DefaultButton(
//                 btnContent: "Register",
//               ),
//               const Gap(24),
//               Align(
//                 alignment: Alignment.center,
//                 child: Text.rich(
//                   TextSpan(
//                     children: [
//                       TextSpan(
//                         recognizer: TapGestureRecognizer()
//                           ..onTap = () => Navigator.pushReplacementNamed(
//                               context, RoutesName.login),
//                         text: 'Sign In',
//                         style: TextStyles.bodyMediumSemiBold.copyWith(
//                           color: Pallete.orangePrimary,
//                           fontSize: getFontSize(14),
//                         ),
//                       ),
//                     ],
//                   ),
//                 ),
//               ),
//             ],
//           ),
//         ),
//       ),
//     );
//   }
// }

import 'package:the_deck/Controller/CustomerController.dart';
import 'package:the_deck/Models/Register.dart';
import 'dart:io';

import 'package:the_deck/Core/Routes/routes_name.dart';
import 'package:the_deck/Core/app_colors.dart';
import 'package:the_deck/Core/font_size.dart';
import 'package:the_deck/Core/response_conf.dart';
import 'package:the_deck/Core/text_styles.dart';
import 'package:the_deck/Presentation/Auth/screens/defaultPw_field.dart';
import 'package:the_deck/Presentation/Auth/screens/default_button.dart';
import 'package:the_deck/Presentation/Auth/screens/default_field.dart';
import 'package:flutter/gestures.dart';
import 'package:flutter/material.dart';
import 'package:gap/gap.dart';
import 'package:image_picker/image_picker.dart';

class SignUpView extends StatefulWidget {
  const SignUpView({Key? key}) : super(key: key);

  @override
  _SignUpViewState createState() => _SignUpViewState();
}

class _SignUpViewState extends State<SignUpView> {
  final ImagePicker _picker = ImagePicker();
  XFile? _image;
  String? _selectedDate;
  final _formKey = GlobalKey<FormState>();

  // Controllers for the text fields
  final TextEditingController _fullnameController = TextEditingController();
  final TextEditingController _usernameController = TextEditingController();
  final TextEditingController _emailController = TextEditingController();
  final TextEditingController _passwordController = TextEditingController();
  final TextEditingController _phoneController = TextEditingController();
  final TextEditingController _addressController = TextEditingController();
  final TextEditingController _genderController = TextEditingController();
  final TextEditingController _dateOfBirthController = TextEditingController();

  @override
  Widget build(BuildContext context) {
    MathUtils.init(context);

    return Scaffold(
      resizeToAvoidBottomInset: false,
      body: Padding(
        padding: EdgeInsets.symmetric(horizontal: getWidth(24)).copyWith(
          top: MediaQuery.of(context).viewPadding.top,
        ),
        child: SingleChildScrollView(
          child: Form(
            key: _formKey,
            child: Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                const Gap(32),
                Text(
                  "Create your new \naccount",
                  style: TextStyles.headingH4SemiBold.copyWith(
                    color: Pallete.neutral100,
                    fontSize: getFontSize(FontSizes.h4),
                  ),
                ),
                const Gap(8),
                Text(
                  "Create an account to start looking for the food \nyou like",
                  style: TextStyles.bodyMediumMedium.copyWith(
                    color: Pallete.neutral60,
                    fontSize: getFontSize(FontSizes.medium),
                  ),
                ),
                const Gap(12),
                DefaultField(
                  hintText: " Full Name",
                  labelText: " Name",
                  controller: _fullnameController,
                ),
                const Gap(14),
                DefaultField(
                  hintText: "Enter Email",
                  labelText: "Email Address",
                  controller: _emailController,
                ),
                const Gap(14),
                DefaultField(
                  hintText: "User Name",
                  labelText: "User Name",
                  controller: _usernameController,
                ),
                const Gap(14),
                DefaultFieldPW(
                  hintText: "Password",
                  labelText: "Password",
                  isPasswordField: true,
                  controller: _passwordController,
                ),
                const Gap(14),
                DefaultField(
                  hintText: "Phone Number",
                  labelText: "Phone Number",
                  controller: _phoneController,
                ),
                const Gap(14),
                DefaultField(
                  hintText: "Address",
                  labelText: "Address",
                  controller: _addressController,
                ),
                const Gap(14),
                const Gap(14),
                Text(
                  "Gender",
                  style: TextStyles.bodyMediumMedium.copyWith(
                    color: Pallete.neutral100,
                    fontSize: getFontSize(14),
                  ),
                ),
                const Gap(8),
                // Ganti dengan DropdownButtonFormField
                DropdownButtonFormField(
                  value: _genderController.text.isNotEmpty
                      ? _genderController.text
                      : null,
                  items: [
                    DropdownMenuItem(
                      value: 'Laki-laki',
                      child: Text('Laki-laki'),
                    ),
                    DropdownMenuItem(
                      value: 'Perempuan',
                      child: Text('Perempuan'),
                    ),
                  ],
                  onChanged: (value) {
                    setState(() {
                      _genderController.text = value.toString();
                    });
                  },
                ),
                const Gap(14),
                GestureDetector(
                  onTap: () async {
                    DateTime? pickedDate = await showDatePicker(
                      context: context,
                      initialDate: DateTime.now(),
                      firstDate: DateTime(1900),
                      lastDate: DateTime(2100),
                    );
                    if (pickedDate != null) {
                      setState(() {
                        _selectedDate = "${pickedDate.toLocal()}".split(' ')[0];
                        _dateOfBirthController.text = _selectedDate!;
                      });
                    }
                  },
                  child: AbsorbPointer(
                    child: DefaultField(
                      hintText: "Date of Birth (YYYY-MM-DD)",
                      labelText: "Date of Birth",
                      controller: _dateOfBirthController,
                    ),
                  ),
                ),
                const Gap(14),
                GestureDetector(
                  onTap: () async {
                    final XFile? pickedFile = await _picker.pickImage(
                      source: ImageSource.gallery,
                    );
                    setState(() {
                      _image = pickedFile;
                    });
                  },
                  child: Container(
                    width: double.infinity,
                    height: getHeight(150),
                    decoration: BoxDecoration(
                      color: Pallete.neutral20,
                      borderRadius: BorderRadius.circular(8),
                      border: Border.all(color: Pallete.neutral40),
                    ),
                    child: _image != null
                        ? Image.file(
                            File(_image!.path),
                            fit: BoxFit.cover,
                          )
                        : Column(
                            mainAxisAlignment: MainAxisAlignment.center,
                            children: [
                              Icon(
                                Icons.camera_alt,
                                color: Pallete.neutral60,
                                size: getSize(50),
                              ),
                              const Gap(8),
                              Text(
                                "Upload Profile Picture",
                                style: TextStyles.bodyMediumMedium.copyWith(
                                  color: Pallete.neutral60,
                                  fontSize: getFontSize(FontSizes.medium),
                                ),
                              ),
                            ],
                          ),
                  ),
                ),
                const Gap(24),
                DefaultButton(
                  btnContent: "Register",
                  function: _register,
                ),
                const Gap(24),
                Align(
                  alignment: Alignment.center,
                  child: Text.rich(
                    TextSpan(
                      children: [
                        TextSpan(
                          recognizer: TapGestureRecognizer()
                            ..onTap = () => Navigator.pushReplacementNamed(
                                context, RoutesName.login),
                          text: 'Sign In',
                          style: TextStyles.bodyMediumSemiBold.copyWith(
                            color: Pallete.orangePrimary,
                            fontSize: getFontSize(14),
                          ),
                        ),
                      ],
                    ),
                  ),
                ),
              ],
            ),
          ),
        ),
      ),
    );
  }

  void _register() {
    if (_formKey.currentState!.validate()) {
      final registerModel = RegisterModel(
        name: _fullnameController.text,
        username: _usernameController.text,
        email: _emailController.text,
        password: _passwordController.text,
        phone: _phoneController.text,
        address: _addressController.text,
        gender: _genderController.text,
        dateOfBirth: _dateOfBirthController.text,
        image: _image?.path ?? '',
      );

      // Pass the model to the controller to handle the registration
      RegisterController().registerUser(registerModel);
    }
  }
}
