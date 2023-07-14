'use client';

import { useState } from 'react';
import styles from './page.module.css';
import { ButtonComponent, InputComponent } from '@/shared/components';

export default function Home() {
  const [counter, setCounter] = useState(0);
  return (
    <main>
      <h1>Counter: {counter}</h1>
      <ButtonComponent title="Click me" onClick={() => setCounter((prev) => prev + 1)} />
      <br />
      <InputComponent type="text" />
    </main>
  );
}
