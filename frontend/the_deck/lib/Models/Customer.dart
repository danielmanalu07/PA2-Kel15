class Customer {
  final String name;
  final String username;
  final String email;
  final String phone;
  final String address;
  final String gender;
  final String dateOfBirth;
  final String image;

  Customer({
    required this.name,
    required this.username,
    required this.email,
    required this.phone,
    required this.address,
    required this.gender,
    required this.dateOfBirth,
    required this.image,
  });

  factory Customer.fromJson(Map<String, dynamic> json) {
    return Customer(
      name: json['name'] ?? '',
      username: json['username'],
      email: json['email'] ?? '',
      phone: json['phone'],
      address: json['address'],
      gender: json['gender'],
      dateOfBirth: json['date_of_birth'],
      image: json['image'],
    );
  }
}
