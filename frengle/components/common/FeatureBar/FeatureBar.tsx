import { FC, ReactNode } from 'react'
import styles from './FeatureBar.module.css'

interface FeatureBarProps {
  className?: string
  title: string
  description?: string
  open?: boolean
  action?: ReactNode
}

const FeatureBar: FC<FeatureBarProps> = ({
  title,
  description,
  className,
  action,
  open,
}) => {
  return (
    <div
      className={[
        styles.FeatureBar,
        open && styles.FeatureBar_open,
        className,
      ].join(' ')}
    >
      <span className={styles.FeatureBar__Title}>{title}</span>
      <span className={styles.FeatureBar__Description}>{description}</span>
      {action && action}
    </div>
  )
}

export default FeatureBar
