class RegisterModel {
  final String name;
  final String username;
  final String email;
  final String password;
  final String phone;
  final String address;
  final String gender;
  final String dateOfBirth;
  final String image;

  RegisterModel({
    required this.name,
    required this.username,
    required this.email,
    required this.password,
    required this.phone,
    required this.address,
    required this.gender,
    required this.dateOfBirth,
    required this.image,
  });

  Map<String, dynamic> toJson() {
    return {
      'fullname': name,
      'username': username,
      'email': email,
      'password': password,
      'phone': phone,
      'address': address,
      'gender': gender,
      'DateOfBirth': dateOfBirth,
      'image': image,
    };
  }
}
