import Link from 'next/link'
import navigation from '../../navigations/navigation'

export default function ButtonLink() {
  return <Link href={navigation.SettingPage.link}>Dashboard</Link>
}