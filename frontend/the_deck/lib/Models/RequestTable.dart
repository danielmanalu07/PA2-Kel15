import 'package:the_deck/Models/Table.dart';

class RequestTable {
  final int id;
  final int tableId;
  final int status;
  final String notes;
  Table? table;

  RequestTable({
    required this.id,
    required this.tableId,
    required this.status,
    required this.notes,
    this.table,
  });

  factory RequestTable.fromJson(Map<String, dynamic> json) {
    return RequestTable(
      id: json['id'] ?? 0,
      tableId: json['table_id'] ?? 0,
      status: json['status'] ?? 0,
      notes: json['notes'] ?? '',
    );
  }
}
