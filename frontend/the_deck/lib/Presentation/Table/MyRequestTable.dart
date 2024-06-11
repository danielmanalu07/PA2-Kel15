import 'package:flutter/material.dart';
import 'package:get/get.dart';
import 'package:the_deck/Controller/CustomerController.dart';

class MyRequestTableScreen extends StatefulWidget {
  const MyRequestTableScreen({Key? key}) : super(key: key);

  @override
  State<MyRequestTableScreen> createState() => _MyRequestTableScreenState();
}

class _MyRequestTableScreenState extends State<MyRequestTableScreen> {
  final RegisterController customerController = Get.put(RegisterController());

  @override
  void initState() {
    super.initState();
    customerController.getMyRequestTable();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: Text('My Request Tables')),
      body: Obx(() {
        if (customerController.isLoading.value) {
          return Center(child: CircularProgressIndicator());
        } else if (customerController.reqTable.isEmpty) {
          return Center(child: Text('No request tables found'));
        } else {
          return ListView.builder(
            itemCount: customerController.reqTable.length,
            itemBuilder: (context, index) {
              final request = customerController.reqTable[index];
              final table = request.table; // Mendapatkan informasi meja terkait
              return ListTile(
                title: Text('Table Number: ${request.tableId}'),
                subtitle: Column(
                  crossAxisAlignment: CrossAxisAlignment.start,
                  children: [
                    Text('Notes: ${request.notes}'),
                    if (table != null)
                      Text(
                          'Capacity: ${table.capacity}'), // Menampilkan kapasitas meja jika tersedia
                  ],
                ),
                trailing: SingleChildScrollView(
                  child: Column(
                    children: [
                      Text(
                        request.status == 1
                            ? 'Approved'
                            : request.status == 0
                                ? 'Waiting for Confirmation'
                                : request.status == 2
                                    ? 'Rejected'
                                    : 'Finished',
                        style: TextStyle(
                          color: request.status == 1
                              ? Colors.green
                              : request.status == 0
                                  ? Colors.orange
                                  : request.status == 2
                                      ? Colors.red
                                      : Colors.greenAccent,
                        ),
                      ),
                      if (request.status == 0)
                        TextButton(
                          onPressed: () => {
                            customerController.CancleRequestTable(request.id),
                          },
                          child: Text(
                            'Cancel',
                            style: TextStyle(color: Colors.red),
                          ),
                        ),
                    ],
                  ),
                ),
              );
            },
          );
        }
      }),
    );
  }
}
