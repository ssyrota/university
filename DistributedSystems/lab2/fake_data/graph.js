class GraphUser {
  constructor(session, user) {
    this.session = session;
    this.user = user;
  }

  async save() {
    const { login, password, cv } = this.user;
    await this.session.run(
      "MERGE (u:User {login: $login}) SET u.password = $password",
      { login, password }
    );
    const cvNode = new GraphCv(this.session, cv, this.user);
    await cvNode.save();

    return;
  }
}
class GraphCv {
  constructor(session, cv, user) {
    this.session = session;
    this.cv = cv;
    this.user = user;
  }

  async save() {
    const { id, hobbies, jobs } = this.cv;
    await this.session.run(
      "MATCH (u:User {login: $login}) MERGE (c:Cv {id: $id}) MERGE (u)-[:HAS_CV]->(c)",
      {
        id,
        login: this.user.login,
      }
    );
    for await (const hobby of hobbies) {
      const hobbyNode = new GraphHobby(this.session, hobby);
      await hobbyNode.save();
      await this.session.run(
        "MATCH (cv:Cv {id: $id}), (h:Hobby {name: $hobbyName}) MERGE (cv)-[:HAS_HOBBY]->(h)",
        { id, hobbyName: hobby.name }
      );
    }
    for await (const job of jobs) {
      const jobNode = new GraphJob(this.session, job);
      await jobNode.save();
      await this.session.run(
        "MATCH (cv:Cv {id: $id}), (j:Job {id: $jobId}) MERGE (cv)-[:HAS_JOB]->(j)",
        { id, jobId: job.id }
      );
    }
    return;
  }
}

class GraphJob {
  constructor(session, job) {
    this.session = session;
    this.job = job;
  }

  async save() {
    const { id, from, to, company, city } = this.job;
    const cityNode = new GraphCity(this.session, city);
    await cityNode.save();
    const companyNode = new GraphCompany(this.session, company);
    await companyNode.save();
    await this.session.run(
      "MATCH (c:Company {name: $companyName}), (ci:City {name: $cityName}) MERGE (j:Job {id: $id}) SET j.from = $from, j.to = $to MERGE (j)-[:IN_CITY]->(ci) MERGE (j)-[:AT_COMPANY]->(c)",
      { id, from, to, companyName: company.name, cityName: city.name }
    );
    return;
  }
}

class GraphCompany {
  constructor(session, company) {
    this.session = session;
    this.company = company;
  }

  async save() {
    const { id, name } = this.company;
    await this.session.run("MERGE (c:Company {name: $name, id:$id})", {
      name,
      id,
    });
    return;
  }
}

class GraphCity {
  constructor(session, city) {
    this.session = session;
    this.city = city;
  }

  async save() {
    const { id, name } = this.city;
    await this.session.run("MERGE (c:City {name: $name, id:$id})", {
      name,
      id,
    });
    return;
  }
}

class GraphHobby {
  constructor(session, hobby) {
    this.session = session;
    this.hobby = hobby;
  }

  async save() {
    const { id, name } = this.hobby;
    await this.session.run("MERGE (h:Hobby {name: $name, id: $id})", {
      name,
      id,
    });
    return;
  }
}
module.exports = {
  GraphUser,
};
