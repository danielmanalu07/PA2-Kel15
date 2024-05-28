import 'package:get/get.dart';
import 'package:the_deck/Controller/CustomerController.dart';
import 'package:the_deck/Core/app_colors.dart';
import 'package:the_deck/Models/Customer.dart';
import 'package:the_deck/Presentation/Auth/screens/default_button.dart';
import 'package:the_deck/Presentation/Base/base.dart';
import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:gap/gap.dart';
import 'package:image_picker/image_picker.dart';
import 'dart:io';

class EditPersonalDataView extends StatefulWidget {
  const EditPersonalDataView({Key? key}) : super(key: key);

  @override
  State<EditPersonalDataView> createState() => _EditPersonalDataViewState();
}

class _EditPersonalDataViewState extends State<EditPersonalDataView> {
  final RegisterController _controller = Get.put(RegisterController());

  final TextEditingController _nameController = TextEditingController();
  final TextEditingController _emailController = TextEditingController();
  final TextEditingController _phoneController = TextEditingController();
  final TextEditingController _usernameController = TextEditingController();
  final TextEditingController _addressController = TextEditingController();
  final TextEditingController _dateOfBirthController = TextEditingController();
    final TextEditingController _genderController = TextEditingController();


  String? _selectedGender;
  File? _selectedImage;

  @override
  void initState() {
    super.initState();
    _controller.getUserProfile().then((_) {
      final customer = _controller.userProfile.value;
      if (customer != null) {
        _nameController.text = customer.name;
        _emailController.text = customer.email;
        _phoneController.text = customer.phone;
        _addressController.text = customer.address;
        _dateOfBirthController.text = customer.dateOfBirth;
        _selectedGender = customer.gender == 'laki laki' ? 'Male' : 'Female';
      }
    });
  }

  Future<void> _selectDate() async {
    DateTime? picked = await showDatePicker(
      context: context,
      initialDate: DateTime.now(),
      firstDate: DateTime(1900),
      lastDate: DateTime.now(),
    );
    if (picked != null) {
      setState(() {
        _dateOfBirthController.text = picked.toIso8601String().split('T').first;
      });
    }
  }

  Future<void> _pickImage(ImageSource source) async {
    final ImagePicker _picker = ImagePicker();
    final XFile? image = await _picker.pickImage(source: source);

    if (image != null) {
      setState(() {
        _selectedImage = File(image.path);
      });
    }
  }

  void _updateProfile() {
    print("Updating profile with:");
    print("Name: ${_nameController.text}");
    print("Date of Birth: ${_dateOfBirthController.text}");

    final updatedCustomer = Customer(
      name: _nameController.text,
      username: _usernameController.text,
      email: _emailController.text,
      phone: _phoneController.text,
      address: _addressController.text,
      image: _selectedImage?.path ?? _controller.userProfile.value?.image ?? '',
      dateOfBirth: _dateOfBirthController.text,
      gender: _selectedGender ?? '',
    );
    _controller.updateUserProfile(updatedCustomer).then((_) {
      // Debugging check to confirm update
      final updatedCustomer = _controller.userProfile.value;
      print("Updated customer profile:");
      print("Name: ${updatedCustomer?.name}");
      print("Date of Birth: ${updatedCustomer?.dateOfBirth}");
    });
  }

  @override
  Widget build(BuildContext context) {
    return Obx(() {
      final customer = _controller.userProfile.value;
      if (customer == null) {
        return Scaffold(
          appBar: buildAppBar(
            buildContext: context,
            screenTitle: "Profile Settings",
            isBackup: false,
          ),
          body: Center(child: CircularProgressIndicator()),
        );
      }
      return Scaffold(
        appBar: buildAppBar(
          buildContext: context,
          screenTitle: "Edit Personal Data",
        ),
        body: Padding(
          padding: EdgeInsets.symmetric(horizontal: 24.0),
          child: SingleChildScrollView(
            child: Column(
              crossAxisAlignment: CrossAxisAlignment.center,
              children: [
                const Gap(24),
                Stack(
                  children: [
                    CircleAvatar(
                      backgroundImage: _selectedImage != null
                          ? FileImage(_selectedImage!)
                          : NetworkImage('http://192.168.30.215:8080/customer/image/${customer.image}') as ImageProvider,
                      radius: 50,
                    ),
                    Positioned(
                      left: 72,
                      bottom: 8,
                      child: GestureDetector(
                        onTap: () {
                          showModalBottomSheet(
                            context: context,
                            builder: (context) => Wrap(
                              children: [
                                ListTile(
                                  leading: Icon(Icons.camera_alt),
                                  title: Text('Camera'),
                                  onTap: () {
                                    _pickImage(ImageSource.camera);
                                    Navigator.pop(context);
                                  },
                                ),
                                ListTile(
                                  leading: Icon(Icons.photo_library),
                                  title: Text('Gallery'),
                                  onTap: () {
                                    _pickImage(ImageSource.gallery);
                                    Navigator.pop(context);
                                  },
                                ),
                              ],
                            ),
                          );
                        },
                        child: Container(
                          width: 32,
                          height: 32,
                          padding: EdgeInsets.all(6),
                          decoration: const BoxDecoration(
                            color: Color(0xFFF5F5FF),
                            shape: BoxShape.circle,
                          ),
                          child: Icon(
                            CupertinoIcons.camera_fill,
                            color: Pallete.orangePrimary,
                            size: 20,
                          ),
                        ),
                      ),
                    ),
                  ],
                ),
                const Gap(24),
                TextField(
                  controller: _nameController,
                  decoration: InputDecoration(labelText: "Full Name"),
                ),
                const Gap(12),
                TextField(
                  controller: _dateOfBirthController,
                  readOnly: true,
                  onTap: _selectDate,
                  decoration: InputDecoration(
                    labelText: "Date of Birth",
                    suffixIcon: Icon(Icons.calendar_today),
                  ),
                ),
                const Gap(12),
                TextField(
                  controller: _phoneController,
                  decoration: InputDecoration(labelText: "Phone"),
                ),
                const Gap(12),
                TextField(
                  controller: _emailController,
                  decoration: InputDecoration(labelText: "Email"),
                ),
                const Gap(12),
                TextField(
                  controller: _addressController,
                  decoration: InputDecoration(labelText: "Address"),
                ),
                const Gap(12),
                TextField(
                  controller: _genderController,
                  decoration: InputDecoration(labelText: "Gender"),
                ),
                const Gap(36),
                DefaultButton(
                  btnContent: "Update Profile",
                  function: _updateProfile,
                ),
              ],
            ),
          ),
        ),
      );
    });
  }
}
