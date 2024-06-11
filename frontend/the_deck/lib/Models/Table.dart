class Table {
  final int id;
  final int number;
  final int capacity;
  final int status;

  Table(
      {required this.id,
      required this.number,
      required this.capacity,
      required this.status});

  factory Table.fromJson(Map<String, dynamic> json) {
    return Table(
      id: json['id'],
      number: json['number'],
      capacity: json['capacity'],
      status: json['status'],
    );
  }
}
