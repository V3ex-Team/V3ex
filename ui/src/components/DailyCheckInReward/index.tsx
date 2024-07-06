import React from 'react';
import styles from './index.module.css';

interface IconProps {
  src: string;
  alt: string;
  className: string;
}

const Icon: React.FC<IconProps> = ({ src, alt, className }) => (
  <img loading="lazy" src={src} alt={alt} className={className} />
);

const DailyCheckIn: React.FC = () => {
  return (
    <section className={styles.container}>
      <div className={styles.header}>
        <Icon
          src="https://cdn.builder.io/api/v1/image/assets/TEMP/a8250ebc45bd6127f6ffa7877085125c72b141ec6b62b95c871c97fde400461a?apiKey=ee1c439a7e5b4920840d177114c3799f&"
          alt=""
          className={styles.icon}
        />
        <Icon
          src="https://cdn.builder.io/api/v1/image/assets/TEMP/4887775fab720c3d26f0636cfe35338287216a82fdb6da2c447a590678a3581c?apiKey=ee1c439a7e5b4920840d177114c3799f&"
          alt=""
          className={styles.icon}
        />
      </div>
      <div className={styles.content}>
        <div className={styles.rewardSection}>
          <div className={styles.rewardInfo}>
            <h2 className={styles.title}>Daily Check-in Reward</h2>
            <img
              loading="lazy"
              src="https://cdn.builder.io/api/v1/image/assets/TEMP/087003798e599646e03f72aa0f190f824634c019e0a03a49808e15bfbcbbccf4?apiKey=ee1c439a7e5b4920840d177114c3799f&"
              alt="Daily reward"
              className={styles.rewardImage}
            />
            <div className={styles.divider} />
            <div className={styles.checkInButton}>
              <div className={styles.buttonSide} />
              <button className={styles.buttonText}>CHECK IN</button>
              <div className={styles.buttonSide} />
            </div>
            <div className={styles.divider} />
          </div>
        </div>
      </div>

      <div className={styles.header}>
        <Icon
          src="https://cdn.builder.io/api/v1/image/assets/TEMP/39e416d959ae72ae68a7479f274fd770b5f56185fedac9992b25aa2dc17792f1?apiKey=ee1c439a7e5b4920840d177114c3799f&"
          alt=""
          className={styles.icon}
        />
        <Icon
          src="https://cdn.builder.io/api/v1/image/assets/TEMP/eebef075674f02e5ce64e82e014d0281b019d938dea8301a28302359af59e09a?apiKey=ee1c439a7e5b4920840d177114c3799f&"
          alt=""
          className={styles.icon}
        />
      </div>
    </section>
  );
};

export default DailyCheckIn;
