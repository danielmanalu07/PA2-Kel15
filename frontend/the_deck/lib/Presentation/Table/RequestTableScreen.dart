import 'package:flutter/material.dart';
import 'package:get/get.dart';
import 'package:the_deck/Controller/CustomerController.dart';
import 'package:the_deck/Controller/TableController.dart';

class RequestTableScreen extends StatefulWidget {
  @override
  _RequestTableScreenState createState() => _RequestTableScreenState();
}

class _RequestTableScreenState extends State<RequestTableScreen> {
  final TableController tableController = Get.put(TableController());
  final TextEditingController notesController = TextEditingController();
  late final RegisterController customerController;
  int? selectedTableId;

  @override
  void initState() {
    super.initState();
    customerController = Get.put(RegisterController());
  }

  void _requestTable() {
    if (selectedTableId == null) {
      Get.snackbar('Error', 'Please select a table');
      return;
    }
    final notes = notesController.text.trim();
    customerController.requestTable(selectedTableId!, notes);
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text('Request Table'),
      ),
      body: Padding(
        padding: const EdgeInsets.all(16.0),
        child: Column(
          children: [
            DropdownButtonFormField<int>(
              decoration: InputDecoration(
                labelText: 'Select Table',
                border: OutlineInputBorder(),
              ),
              items: tableController.tables
                  .map((table) => DropdownMenuItem<int>(
                        value: table.id,
                        child: Text(
                            'Table ${table.number} (Capacity: ${table.capacity})'),
                      ))
                  .toList(),
              onChanged: (value) {
                setState(() {
                  selectedTableId = value;
                });
              },
              value: selectedTableId,
            ),
            SizedBox(height: 16),
            TextField(
              controller: notesController,
              decoration: InputDecoration(
                labelText: 'Notes',
                border: OutlineInputBorder(),
              ),
              maxLines: 3,
            ),
            SizedBox(height: 16),
            ElevatedButton(
              onPressed: _requestTable,
              child: Text('Request Table'),
            ),
          ],
        ),
      ),
    );
  }
}
