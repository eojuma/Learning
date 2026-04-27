var user = {
  firstName: "John",
  lastName: "K'Omollo",
  role: "admin",
  loginCount: 4,
  facebookLoggin: true,
  courseList: [],
  buyCourse: function (courseName) {
    this.courseList.push(courseName);
  },
  getCourseCount: function () {
    return `${this.firstName} is enrolled in a total of ${this.courseList.length} courses`;
  },

  info: function () {
    return `
${this.firstName};
${this.lastName};
${this.role};
${this.loginCount};
${this.facebookLoggin};
${this.courseList}`;
  },
};
user.buyCourse("ReactJs");
console.log(user.getCourseCount());
console.log(user.info());
