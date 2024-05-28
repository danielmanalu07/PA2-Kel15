class Table {
  final int id;
  final String number;
  final String capacity;

  Table({required this.id, required this.number, required this.capacity});

  factory Table.fromJson(Map<String, dynamic> json) {
    return Table(
      id: json['id'],
      number: json['number'],
      capacity: json['capacity'],
    );
  }
}
