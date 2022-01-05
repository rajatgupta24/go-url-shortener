import styles from '../styles/Home.module.css'

const HelloWorld = () => {
  return (
    <div className={styles.container}>
      <main className={styles.main}>
        <h1 className={styles.title}>
          Welcome to <a href="https://github.com">Github</a>
        </h1>
      </main>
    </div>
  )
}

export default HelloWorld;
