import React from 'react';
import './input.component.css';

interface InputComponentProps {
  type: 'text' | 'password' | 'email' | 'number';
}

export default function InputComponent(props: InputComponentProps) {
  return <input type={props.type}></input>;
}
