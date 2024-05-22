import 'package:flutter/material.dart';

class ReadOnlyField extends StatelessWidget {
  final String labelText;
  final String value;

  const ReadOnlyField({
    Key? key,
    required this.labelText,
    required this.value,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return TextFormField(
      readOnly: true,
      initialValue: value,
      decoration: InputDecoration(
        labelText: labelText,
        border: OutlineInputBorder(),
      ),
    );
  }
}
